import axios from "axios";
const baseURL = "http://localhost:8080/api/datas/";
class ApiServices {
  getAllData() {
    return axios.get(`${baseURL}`);
  }
  postData(data) {
    return axios.post(`${baseURL}`, data);
  }
  removeData(id) {
    return axios.delete(`${baseURL}${id}`);
  }
  getData(id) {
    return axios.get(`${baseURL}${id}`);
  }

  updateData(id, title, CreatedAt) {
    return axios.put(`${baseURL}${id}`, { id, title, CreatedAt });
  }
}
export default new ApiServices();
