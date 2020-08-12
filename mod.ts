import {
  Application,
  Router,
  Context,
  MongoClient,
  Result,
} from "./deps.ts";

import {
  availableSchools,
  MongoSchoolRepository,
  createSchool,
  detailedSchoolInformation,
  availableCanteens,
  createCanteen,
  detailedCanteenInformation,
  availableDishes,
  detailedDishInformation,
  availableMenus,
  createMenu,
  detailedMenuInformation,
} from "./controllers/controllers.ts";

import { School } from "./models/models.ts";

import {
  BadRequest,
  InternalServerError,
  NotFound,
} from "./views/views.ts";

function respond(
  ctx: Context,
  view: Result<Object, Object>,
  successStatusCode: number,
): void {
  if (view.isErr()) {
    respondWithError(ctx, view.unwrapErr());
  } else {
    ctx.response.status = successStatusCode;
    ctx.response.body = view.unwrap();
  }
}

function respondWithError(
  ctx: Context,
  error: BadRequest | InternalServerError | NotFound,
): void {
  console.warn(`respondWithError called with: ${error}`);

  if (error instanceof BadRequest) {
    ctx.response.status = 400;
    ctx.response.body = error;
  } else if (error instanceof NotFound) {
    ctx.response.status = 404;
  } else {
    ctx.response.status = 500;
  }
}

async function authorizationMiddleware(
  ctx: Context,
  next: () => Promise<void>,
) {
  if (ctx.request.method === "POST") {
    const authKey = (ctx.request.headers.get("Authorization") ||
      "Basic ").replace("Basic ", "");
    if (await isAuthorizedForPostRequests(authKey, ipeaUrl)) {
      await next();
    } else {
      console.warn(
        "POST Request failed authorization check... Request Headers:",
      );
      console.warn(ctx.request.headers);
      ctx.response.status = 401;
    }
  } else {
    await next();
  }
}

async function isAuthorizedForPostRequests(
  authKey: string,
  ipeaUrl: string,
): Promise<boolean> {
  const response = await fetch(`${ipeaUrl}/${authKey}`);
  return response.ok;
}

const client = new MongoClient();

const ipeaUrl = Deno.env.get("IPEA_URL") || "undefined";

client.connectWithUri(
  Deno.env.get("MONGO_DB_CONNECTION_STRING") || "undefined",
);

const db = client.database("ipp-ementa");

const schoolCollection = db.collection<School>("school");

const schoolRepository = new MongoSchoolRepository(schoolCollection);

const router = new Router();

router.get("/schools", async (ctx) => {
  const view = await availableSchools(
    schoolRepository,
  );

  respond(ctx, view, 200);
});

router.post("/schools", async (ctx) => {
  const schoolToCreate = await ctx.request.body().value;

  const view = await createSchool(
    schoolRepository,
    schoolToCreate,
  );

  respond(ctx, view, 201);
});

router.get("/schools/:id", async (ctx) => {
  const schoolId = ctx.params.id || "undefined";

  const view = await detailedSchoolInformation(
    schoolRepository,
    schoolId,
  );

  respond(ctx, view, 200);
});

router.get("/schools/:id/canteens", async (ctx) => {
  const schoolId = ctx.params.id || "undefined";

  const view = await availableCanteens(
    schoolRepository,
    schoolId,
  );

  respond(ctx, view, 200);
});

router.post("/schools/:id/canteens", async (ctx) => {
  const schoolId = ctx.params.id || "undefined";

  const canteenToCreate = await ctx.request.body().value;

  const view = await createCanteen(
    schoolRepository,
    schoolId,
    canteenToCreate,
  );

  respond(ctx, view, 201);
});

router.get("/schools/:id1/canteens/:id2", async (ctx) => {
  const schoolId = ctx.params.id1 || "undefined";

  const canteenId = ctx.params.id2 || "undefined";

  const view = await detailedCanteenInformation(
    schoolRepository,
    schoolId,
    canteenId,
  );

  respond(ctx, view, 200);
});

router.get("/schools/:id1/canteens/:id2/menus", async (ctx) => {
  const schoolId = ctx.params.id1 || "undefined";

  const canteenId = ctx.params.id2 || "undefined";

  const view = await availableMenus(
    schoolRepository,
    schoolId,
    canteenId,
  );

  respond(ctx, view, 200);
});

router.post("/schools/:id1/canteens/:id2/menus", async (ctx) => {
  const schoolId = ctx.params.id1 || "undefined";

  const canteenId = ctx.params.id2 || "undefined";

  const menuToCreate = await ctx.request.body().value;

  const view = await createMenu(
    schoolRepository,
    schoolId,
    canteenId,
    menuToCreate,
  );

  respond(ctx, view, 200);
});

router.get("/schools/:id1/canteens/:id2/menus/:id3", async (ctx) => {
  const schoolId = ctx.params.id1 || "undefined";

  const canteenId = ctx.params.id2 || "undefined";

  const menuId = ctx.params.id3 || "undefined";

  const view = await detailedMenuInformation(
    schoolRepository,
    schoolId,
    canteenId,
    menuId,
  );

  respond(ctx, view, 200);
});

router.get("/schools/:id1/canteens/:id2/menus/:id3/dishes", async (ctx) => {
  const schoolId = ctx.params.id1 || "undefined";

  const canteenId = ctx.params.id2 || "undefined";

  const menuId = ctx.params.id3 || "undefined";

  const view = await availableDishes(
    schoolRepository,
    schoolId,
    canteenId,
    menuId,
  );

  respond(ctx, view, 200);
});

router.get(
  "/schools/:id1/canteens/:id2/menus/:id3/dishes/:id4",
  async (ctx) => {
    const schoolId = ctx.params.id1 || "undefined";

    const canteenId = ctx.params.id2 || "undefined";

    const menuId = ctx.params.id3 || "undefined";

    const dishId = ctx.params.id4 || "undefined";

    const view = await detailedDishInformation(
      schoolRepository,
      schoolId,
      canteenId,
      menuId,
      dishId,
    );

    respond(ctx, view, 200);
  },
);

const app = new Application();

app.use(authorizationMiddleware);

app.use(router.routes());

app.use(router.allowedMethods());

await app.listen(`:${Deno.env.get("PORT")}`);
