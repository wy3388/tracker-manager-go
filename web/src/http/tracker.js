import {get} from "@/http/index.js";

class TrackerService {
    static async getTracker() {
        return get(`tracker`)
    }
}

export {
    TrackerService
}