import { AuthService } from "./common/auth.service.js";
import { ListService } from "./common/list.service.js";
import { ClientDataService } from "./common/client-data.service.js";

(() => {
  const APP = {
    init: () => {
      console.log("App initialized!");
      ClientDataService.getClientDataConfiguration().then((data) => {
        console.log(data);
      });
    },

    ClientDataService,
    AuthService,
    ListService,
  };

  window.onload = () => {
    APP.init();
    window["APP"] = APP;
  };
})();
