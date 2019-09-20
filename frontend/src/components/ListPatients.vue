<template>
    <div>
        <h1>Patient List</h1>

        <!-- Patient Table -->
        <b-table hover :items="patients" :fields="fields">
            <template v-slot:cell(action)="data">
                <b-button size="sm" class="mx-1" v-b-modal="'edit'" v-on:click="updatePlaceholder(data.item)">Edit</b-button>
                <b-button size="sm" class="mx-1" v-on:click="deletePatient(data.item.id)"> Delete </b-button>
            </template>
        </b-table>
    
        <!-- Add Patient -->
        <b-table-simple hover> 
            <b-tbody>
                <b-tr>
                    <b-td colspan="2">
                        <div class="input-field">
                        <label>First Name</label>
                        <input placeholder="" v-model="input.firstName" type="text">
                        </div>
                    </b-td>
                    <b-td>
                        <div class="input-field">
                        <label>Last Name</label>
                        <input placeholder="" v-model="input.lastName" type="text">
                        </div>
                    </b-td>
                    <b-td>
                        <div class="input-field">
                        <label>Diagnosis</label>
                        <input placeholder="" v-model="input.diagnosis" type="text">
                        </div>
                    </b-td>
                    <b-td>
                        <div class="input-field">
                        <label>Physician</label>
                        <input placeholder="" v-model="input.physician" type="text">
                        </div>
                    </b-td>
                    <b-td>
                        <div class="input-field">
                        <label>Date of Visit</label>
                        <input :placeholder= "input.dov" v-model="input.dov" type="text">
                        </div>
                    </b-td>
                    <b-td>
                        <b-button class="button" size="sm" v-on:click="addPatient()"> Add </b-button>
                    </b-td>
                </b-tr>
            </b-tbody>
        </b-table-simple>

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
        
        <!-- Error Handling -->
        <b-alert v-model="showAlert" variant="danger"> Please fill in all fields </b-alert>

    </div>
</template>

<style scoped>
    .input-field input[type=text] {
        text-align: center;
    }

    .button {
        display: flex;
        justify-content: center;
        align-items: center;
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
                fields: [
                    {
                        key: "id",
                        label: "Medical #",
                        sortable: true
                    },
                    {
                        key: "firstName",
                        sortable: true
                    },
                    {
                        key: "lastName",
                        sortable: true
                    },
                    {
                        key: "diagnosis",
                        sortable: true
                    },
                    {
                        key: "physician",
                        sortable: true
                    },
                    {
                        key: "dov",
                        label: "Date of Visit",
                        sortable: true
                    },
                    {
                        key: 'action',
                        sortable: false
                    }
                ],
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
                },
                showAlert: false
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
                let input = Object.values(this.input);

                // check if each field is not empty
                for (let value of input) {
                    if (value == "") {
                        this.showAlert = true;
                        return;
                    }
                }
                this.showAlert = false;

                if (!this.showAlert) {
                    APIService.createPatient(this.input).then(res => {
                        this.getPatients();
                    });
                }
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