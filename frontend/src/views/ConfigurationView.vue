<template>
<div>
  <h1>Configurazione</h1>
  <form>
    <table>
      <tr>
        <h2>Motore</h2>
      </tr>
      <tr>
      <td>
    <label for="degreePerSteprequest">Gradi per passo di motore</label></td><td>
    <input id="degreePerSteprequest" type="number" step="0.1" v-model="configuration.hardware.motorDegreePerStep" /></td></tr><tr><td>
    <label for="millisecsPerStep">Millisecondi per passo di motore</label></td><td>
    <input id="millisecsPerStep" type="number" v-model="configuration.hardware.waitForStep" /></td></tr><tr><td>
    <label for="ratioMotor">Rapporto trasmissione albero motore/albero tornio</label></td><td>
    <input id="ratioMotor" type="number" v-model="configuration.hardware.gearRatio" /></td></tr>
    <tr>
        <h2>Camera</h2>
      </tr>
      <tr>
      <td>
    <label for="height">Altezza immagine px</label></td><td>
    <input id="height" type="number" v-model="configuration.hardware.camera.height" /></td></tr><tr><td>
    <label for="width">Larghezza immagine px</label></td><td>
    <input id="width" type="number" v-model="configuration.hardware.camera.width" /></td></tr><tr><td>
      <label for="brightness">Luminosità</label></td><td>
    <input id="brightness" type="number" v-model="configuration.hardware.camera.brightness" /></td></tr><tr><td>
      <label for="contrast">Contrasto</label></td><td>
    <input id="contrast" type="number" v-model="configuration.hardware.camera.contrast" /></td></tr><tr><td>
      <label for="sharpness">Definizione</label></td><td>
    <input id="sharpness" type="number" v-model="configuration.hardware.camera.sharpness" /></td></tr><tr><td>
     <label for="num-photo">Numero foto per ciclo</label></td><td>
    <input id="num-photo" type="number" v-model="configuration.hardware.camera.numPhotosPerProcess" /></td></tr>
      <tr>
        <h2>Server</h2>
      </tr>
      <tr>
      <td>
    <label for="imagePath">Path immagini</label></td><td>
    <input id="imagePath" type="text" v-model="configuration.server.photoDirectory" /></td></tr><tr><td>
    <label for="serverPort">Porta server (richiede restart)</label></td><td>
    <input id="serverPort" type="text" v-model="configuration.server.port" /></td></tr><tr>
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
          gearRatio:0,
          camera:{
            height:0,
            width: 0,
            brightness: 0,
            sharpness: 0,
            contrast: 0,
          }
        },
        server: {
          photoDirectory:"",
          distributionDirectory:"",
          port: 3333
        }
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
    })
  }
}
</script>

<style>

</style>