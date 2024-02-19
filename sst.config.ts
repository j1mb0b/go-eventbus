import { SSTConfig } from "sst";
import { Api } from "sst/constructs";

export default {
  config(_input) {
    return {
      name: "go-eventbus",
      region: "eu-west-2",
    };
  },
  stacks(app) {
    app.setDefaultFunctionProps({
      runtime: "go",
    });
    app.stack(function Stack({ stack }) {
      const api = new Api(stack, "api", {
        routes: {
          "GET /": "functions/lambda/main.go",
          "POST /orders": "functions/orders/create_order.go",
        },
      });
      stack.addOutputs({
        ApiEndpoint: api.url,
      });
    });
  },
} satisfies SSTConfig;
