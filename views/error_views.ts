export class Error {}

export class BadRequest extends Error {
  message: string;

  constructor(message: string) {
    super();
    this.message = message;
  }

  toString() {
    return "BadRequest";
  }
}

export class NotFound extends Error {
  toString() {
    return "NotFound";
  }
}

export class InternalServerError extends Error {
  toString() {
    return "InternalServerError";
  }
}
