import {PluginDefinition} from '@yaakapp/api';

export const plugin: PluginDefinition = {
    httpRequestActions: [{
        key: 'export-curl',
        label: 'Copy as Curl',
        icon: 'copy',
        async onSelect(ctx, args) {
            console.log("Hello World!");
        },
    }],
};
