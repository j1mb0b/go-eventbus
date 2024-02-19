import {EventBus, Api, StackContext} from "@serverless-stack/resources";

export function Stack({stack}: StackContext) {
    // 1. Create event bus
    const bus = new EventBus(stack, "EventBus", {})

    // 2. Add set up event routing rules
    bus.addRules(stack, {
        "order_created": {
            pattern: {
                detailType: ["order.created"],
            },
            targets: {
                function: "functions/logging/log_events.go"
            }
        }
    })

    const api = new Api(stack, "Api", {
        routes: {
            "POST /order": "functions/orders/create_order.go",
        },
        defaults: {
            functions: {
                permissions: [bus],
                envrionment: {
                    EVENT_BUS_NAME: bus.eventBusName
                }
            }
        }
    });

	// 3. Output the event bus name
    stack.addOutputs({
        ApiEndpoint: api.url,
        BusName: bus.eventBusName
    });
}