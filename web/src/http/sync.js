import {get, post} from "@/http/index.js";

class SyncService {
    static async sync() {
        return post('sync/syncTracker')
    }

    static async deleteByIds(params) {
        return post('sync/deleteByIds', params)
    }

    static async clear() {
        return get('sync/clear')
    }
}

export {
    SyncService
}