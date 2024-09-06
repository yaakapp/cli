import { PluginDefinition } from "@yaakapp/api";

export const plugin: PluginDefinition = {
    httpRequestActions: [
        {
            key: "test-example-plugin",
            label: "Test Example Plugin",
            icon: "cake",
            async onSelect(ctx, args) {
                ctx.toast.show({ variant: "success", message: "It works, have some cake!" });
            },
        },
    ],
};
