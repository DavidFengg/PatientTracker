import axios from "axios";

const API_URL = 'http://localhost:8080/patient';

export default {
    getPatients() {
        return axios.get(API_URL).then(res => {
            return res.data;
        });   
    },

    getPatient(id) {
        return axios.get(API_URL + "/" + id).then(res => {
            return res.data;
        })
    },

    createPatient(fName, lName, diag, phys, date) {
        return axios.post(API_URL, {
            firstName: fName,
            lastName: lName,
            diagnosis: diag,
            physician: phys,
            dov: date
        }).then(res => {
            console.log(res.data);
        }).catch(error => console.log(error));
    },

    updatePatient(data) {
        return axios.put(API_URL + "/" + data.id, {
            firstName: data.firstName,
            lastName: data.lastName,
            diagnosis: data.diagnosis,
            physician: data.physician,
            dov: data.dov
        }).then(res => {
            console.log(res.data);
        }).catch(error => console.log(data.id, error));
    },

    deletePatient(id) {
        return axios.delete(API_URL + "/" + id).then(res => {
            console.log(res.data);
        }).catch(error => console.log(error));
    }
}