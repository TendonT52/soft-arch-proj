window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    // urls: [
    //   swagger-ui/v1
    //   { url: "./v1/auth.swagger.json", name: "authService" },
    //   { url: "./v1/user.swagger.json", name: "userService" },
    // ],
    dom_id: "#swagger-ui",
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: "StandaloneLayout",
    url: "sofe-arch-prog.swagger.json",
  });

  //</editor-fold>
};
