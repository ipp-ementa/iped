import { Application, Router, Context } from "https://deno.land/x/oak/mod.ts";

import { MongoClient, Result } from "./deps.ts";

import {
  availableSchools,
  MongoSchoolRepository,
  createSchool,
  detailedSchoolInformation,
  availableCanteens,
  createCanteen,
  detailedCanteenInformation,
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

const client = new MongoClient();

client.connectWithUri(
  "",
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


const app = new Application();

app.use(router.routes());

app.use(router.allowedMethods());

app.use((ctx) => {
  ctx.response.body = "Hello world!";
});

await app.listen("127.0.0.1:8000");
