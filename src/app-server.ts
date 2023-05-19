import { singleton } from "tsyringe";
import * as http from "http";
import express from "express";
import helmet from "helmet";
import { ControllersIndex } from "./controllers/controllers.index";
import {ConfigService} from "./services/config.service";

@singleton()
export class AppServer {
  private app: express.Application;

  constructor(
      private readonly configService: ConfigService,
      private readonly controllersIndex: ControllersIndex
  ) {
    this.app = express();
  }

  public async start(): Promise<http.Server> {

    this.app.use(helmet());

    this.app.use(express.json());
    this.app.use(express.urlencoded({ extended: true }));

    this.app.use("/api", this.controllersIndex.apiRoutes);
    this.app.use("/", this.controllersIndex.uiRoutes);

    return this.app.listen(this.configService.AppPort);
  }
}
