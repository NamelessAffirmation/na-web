import { Router } from "express";

export class BaseController {
  private readonly _router: Router;

  public get router(): Router {
    return this._router;
  }

  constructor() {
    this._router = Router();
  }
}
