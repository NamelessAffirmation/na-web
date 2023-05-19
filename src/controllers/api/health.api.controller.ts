import { BaseController } from "../base.controller";
import { NextFunction, Request, Response } from "express";
import { singleton } from "tsyringe";

@singleton()
export class HealthAPIController extends BaseController {
  constructor() {
    super();
    this.buildRoutes();
  }

  private health(_req: Request, res: Response, next: NextFunction): void {
    try {
      res.status(200).send({
        status: "ok"
      });
    } catch (error: any) {
      next(error);
    }
  }

  private buildRoutes() {
    this.router.get("/", this.health.bind(this));
  }
}
