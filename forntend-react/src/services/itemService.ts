import axios from "axios"

const API_URL = "http://127.0.0.1:8082"

export const itemService = {
    addItem: async (item: any) => {
        const response = await axios.post(`${API_URL}/items`, item)
        return response.data
    },
    getAllItems: async () => {
        const response = await axios.get(`${API_URL}/items`)
        return response.data
    },
    getSummary: async () => {
        const response = await axios.get(`${API_URL}/items/all/summary`)
        return response.data
    },
}