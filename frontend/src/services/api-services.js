import axios from "axios";

const baseURL = "http://localhost:8080/api/datas/";
class ApiServices {
  getAllData() {
    return axios.get(baseURL);
  }
  postData(data) {
    return axios.post(baseURL, data);
  }
}
export default new ApiServices();
