import { singleton } from "tsyringe";
import { Router } from "express";
import { HealthAPIController } from "./api/health.api.controller";

@singleton()
export class ControllersIndex {
  private readonly _apiRoutes: Router;
  private readonly _uiRoutes: Router;

  public get apiRoutes(): Router {
    return this._apiRoutes;
  }

  public get uiRoutes(): Router {
    return this._uiRoutes;
  }

  constructor(private readonly healthAPIController: HealthAPIController) {
    this._apiRoutes = Router();
    this._uiRoutes = Router();
    this.buildRoutes();
  }

  private buildRoutes() {
    this.buildAPIRoutes();
    this.buildUIRoutes();
  }

  private buildAPIRoutes(): void {
    this._apiRoutes.use("/health", this.healthAPIController.router);
  }

  private buildUIRoutes(): void {}
}
