import { describe, expect, test } from 'vitest';
import {plugin} from "./index";

describe('Example Plugin', () => {
    test('Exports plugin object', () => {
        expect(plugin).toBeTypeOf('object');
    });
});
