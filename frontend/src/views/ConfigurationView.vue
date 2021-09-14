<template>
<div>
  <h1>Configurazione</h1>
  <form>
    <table>
      <tr>
      <td>
    <label for="degreePerSteprequest">Gradi per passo di motore</label></td><td>
    <input id="degreePerSteprequest" type="number" step="0.1" v-model="configuration.hardware.motorDegreePerStep" /></td></tr><tr><td>
    <label for="millisecsPerStep">Millisecondi per passo di motore</label></td><td>
    <input id="millisecsPerStep" type="number" v-model="configuration.hardware.waitForStep" /></td></tr><tr><td>
    <label for="ratioMotor">Rapporto trasmissione albero motore/albero tornio</label></td><td>
    <input id="ratioMotor" type="number" v-model="configuration.hardware.gearRatio" /></td></tr><tr><td>
    <label for="imagePath">Path immagini</label></td><td>
    <input id="imagePath" type="text" v-model="configuration.photoDirectory" /></td></tr><tr><td>
    <label for="distPath">Path folder distribuzione</label></td><td>
    <input id="distPath" type="text" v-model="configuration.distributionDirectory" /></td></tr><tr>
      <td></td><td><button type="button" @click="updateConfiguration">Aggiorna</button></td></tr>
    </table>
  </form>
</div>
</template>

<script>
export default {
  name: "ConfigurationView",
  data() {
    return {
      message:"",
      configuration: {
        hardware: {
          motorDegreePerStep:0,
          waitForStep:0,
          gearRatio:0
        },
        photoDirectory:"",
        distributionDirectory:""
      }
    }
  },
  methods:{
    updateConfiguration(){
      fetch('/api/configuration/' , {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.configuration),
      })
        .then((a) => {
          return a.json()
        })
        .then((data) => {
          this.message = data
        })
    }
    
  },
  mounted(){
    fetch('/api/configuration')
    .then(data=> {
      return data.json()
    })
    .then(config=>{
      this.configuration=config
      console.log(this.configuration)
    })
  }
}
</script>

<style>

</style>