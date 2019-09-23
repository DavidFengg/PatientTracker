<template>
  <div class="home">
    <PatientList/>
  </div>
</template>

<script>
// @ is an alias to /src
import PatientList from '@/components/ListPatients.vue'

export default {
  name: 'home',
  components: {
    PatientList
  }
}
</script>
