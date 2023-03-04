import {del, get, put, post} from "@/http/index"

class SourceService {
    static async getById(id) {
        return get(`source/${id}`)
    }

    static async updateById(params) {
        return put(`source`, params)
    }

    static async deleteById(id) {
        return del(`source/${id}`)
    }

    static async add(params) {
        return post('source', params)
    }
}

export {
    SourceService
}