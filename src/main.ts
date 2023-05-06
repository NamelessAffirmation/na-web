import "reflect-metadata";
import { AppServer } from "./app-server";
import { container } from "tsyringe";
import { AddressInfo } from "net";

const bootstrap = async () => {
  const appServer = container.resolve(AppServer);
  const server = await appServer.start();
  console.log(
    `Server listening on port: [${(server.address() as AddressInfo).port}]`
  );
};

bootstrap().catch((error) => {
  console.error(error.message);
});
