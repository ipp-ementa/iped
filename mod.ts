import { Application, Router } from "https://deno.land/x/oak/mod.ts";

const router = new Router();

router.get("/sample", (ctx) => {
  ctx.response.body = "sample";
});

router.get("/sample/:id", (ctx) => {
  ctx.response.body = `sample with id: ${ctx.params.id}`;
});

const app = new Application();

app.use(router.routes());

app.use(router.allowedMethods());

app.use((ctx) => {
  ctx.response.body = "Hello world!";
});

await app.listen("127.0.0.1:8000");
