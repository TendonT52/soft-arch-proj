window.onload = function () {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    urls: [
      { url: "post-service.swagger.json", name: "Post service" },
      { url: "user-service.swagger.json", name: "User service" },
      { url: "report-service.swagger.json", name: "Report service" },
      { url: "review-service.swagger.json", name: "Review service" },
    ],
    dom_id: "#swagger-ui",
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: "StandaloneLayout",
  });
  //</editor-fold>
};
