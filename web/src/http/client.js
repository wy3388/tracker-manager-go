import {del, get, post, put} from "@/http/index.js"

class ClientService {
    static async add(params) {
        return post('client', params)
    }

    static async getById(id) {
        return get(`client/${id}`)
    }

    static async updateById(params) {
        return put(`client`, params)
    }

    static async deleteById(id) {
        return del(`client/${id}`)
    }
}

export {
    ClientService,
}