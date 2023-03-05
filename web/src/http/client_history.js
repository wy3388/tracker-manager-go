import {post} from "@/http/index.js"

class ClientHistoryService {
    static async deleteByIds(ids) {
        return post('clientHistory/deleteByIds', ids)
    }
}

export {
    ClientHistoryService
}