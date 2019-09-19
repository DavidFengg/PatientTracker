<template>
    <div>
        <h1>Patient List</h1>

        <table class="table table-hover">
            <thead>
                <tr>
                    <th>Medical #</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Diagnosis</th>
                    <th>Physician</th>
                    <th>Date of Visit</th>
                    <th>Action</th>
                </tr>
            </thead>

            <tbody>
                <!-- Display Patients -->
                <tr v-for="patient in patients">
                    <th> {{ patient.id }} </th>
                    <th> {{ patient.firstName }} </th>
                    <th> {{ patient.lastName }} </th>
                    <td> {{ patient.diagnosis }} </td>
                    <td> {{ patient.physician }} </td>
                    <td> {{ patient.dov }} </td>
                    <td>
                        <b-button size="sm" class="mx-1" v-b-modal="'edit'" v-on:click="updatePlaceholder(patient)">Edit</b-button>
                        <b-button size="sm" class="mx-1" v-on:click="deletePatient(patient.id)"> Delete </b-button>
                    </td>
                </tr>

                <!-- Add Patient -->
                <tr>
                    <td colspan="2">
                        <div class="input-field">
                        <label>First Name</label>
                        <input placeholder="" v-model="input.firstName" type="text">
                        </div>
                    </td>
                    <td>
                        <div class="input-field">
                        <label>Last Name</label>
                        <input placeholder="" v-model="input.lastName" type="text">
                        </div>
                    </td>
                    <td>
                        <div class="input-field">
                        <label>Diagnosis</label>
                        <input placeholder="" v-model="input.diagnosis" type="text">
                        </div>
                    </td>
                    <td>
                        <div class="input-field">
                        <label>Physician</label>
                        <input placeholder="" v-model="input.physician" type="text">
                        </div>
                    </td>
                    <td>
                        <div class="input-field">
                        <label>Date of Visit</label>
                        <input :placeholder= "input.dov" v-model="input.dov" type="text">
                        </div>
                    </td>
                    <td>
                        <b-button size="sm" v-on:click="addPatient()"> Add </b-button>
                    </td>
                </tr>
            </tbody>

        </table>

        <!-- Edit Modal -->
        <b-modal id="edit" title="Edit Patient Data">
            <b-form-group label="First Name">
                <b-form-input
                    v-model="edit.firstName"
                    placeholder= "">
                </b-form-input>
            </b-form-group>

            <b-form-group label="Last Name">
                <b-form-input
                    v-model="edit.lastName"
                    placeholder= "">
                </b-form-input>
            </b-form-group>

            <b-form-group label="Diagnosis">
                <b-form-input
                    v-model="edit.diagnosis"
                    placeholder= "">
                </b-form-input>
            </b-form-group>

            <b-form-group label="Physician">
                <b-form-input
                    v-model="edit.physician"
                    placeholder= "">
                </b-form-input>
            </b-form-group>

            <b-form-group label="Date of Visit">
                <b-form-input
                    v-model="edit.dov"
                    placeholder= "">
                </b-form-input>
            </b-form-group>

        <b-button class="btn btn-primary" v-on:click="updatePatient()">Confirm</b-button>

        </b-modal>

    </div>

</template>

<style scoped>
    .input-field input[type=text] {
        text-align: center;
    }
</style>

<script>
    import APIService from '../APIService';

    let today = new Date();
    let date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();

    export default {
        name: 'Patients',

        components: {
        },

        data() {
            return {
                patients: [],
                edit: {
                    id: 0,
                    firstName: "",
                    lastName: "",
                    diagnosis: "",
                    physician: "",
                    dov: ""
                },
                input: {
                    firstName: "",
                    lastName: "",
                    diagnosis: "",
                    physician: "",
                    dov: date
                }
            };
        },

        methods: {
            getPatients() {
                APIService.getPatients().then(data => {
                    this.patients = data;
                }).catch(error => console.log(error));
            },

            getPatient(id) {
            },

            addPatient() {
                APIService.createPatient(this.input).then(res => {
                    this.getPatients();
                });
            },

            updatePlaceholder(patient) {
                // copy patient object data
                let copyData = JSON.parse(JSON.stringify(patient));
                this.edit = copyData;
            },

            updatePatient() {
                APIService.updatePatient(this.edit).then(res => {
                    this.getPatients();
                })
            },
            
            deletePatient(id) {
                APIService.deletePatient(id).then(res => {
                    this.getPatients();
                });
            }
        },

        mounted() {
            this.getPatients();
        }
    }
</script>