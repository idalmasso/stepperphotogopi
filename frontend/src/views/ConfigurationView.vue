<template>
  <div>
    <div class="error-text" v-if="error">{{ error }}</div>
    <h1>Configurazione</h1>
    <form>
      <table>
        <tr>
          <h2>Motore</h2>
        </tr>
        <tr>
          <td>
            <label for="degreePerSteprequest">Gradi per passo di motore</label>
          </td>
          <td>
            <input
              id="degreePerSteprequest"
              type="number"
              step="0.1"
              v-model="configuration.hardware.motorDegreePerStep"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="millisecsPerStep"
              >Millisecondi per passo di motore</label
            >
          </td>
          <td>
            <input
              id="millisecsPerStep"
              type="number"
              v-model="configuration.hardware.waitForStep"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="ratioMotor"
              >Rapporto trasmissione albero motore/albero tornio</label
            >
          </td>
          <td>
            <input
              id="ratioMotor"
              type="number"
              v-model="configuration.hardware.gearRatio"
            />
          </td>
        </tr>
        <tr>
          <h2>Camera</h2>
        </tr>
        <tr>
          <td>
            <label for="height">Altezza immagine px</label>
          </td>
          <td>
            <input
              id="height"
              type="number"
              v-model="configuration.hardware.camera.height"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="width">Larghezza immagine px</label>
          </td>
          <td>
            <input
              id="width"
              type="number"
              v-model="configuration.hardware.camera.width"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="brightness">Luminosità</label>
          </td>
          <td>
            <input
              id="brightness"
              type="number"
              v-model="configuration.hardware.camera.brightness"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="contrast">Contrasto</label>
          </td>
          <td>
            <input
              id="contrast"
              type="number"
              v-model="configuration.hardware.camera.contrast"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="sharpness">Definizione</label>
          </td>
          <td>
            <input
              id="sharpness"
              type="number"
              v-model="configuration.hardware.camera.sharpness"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="saturation">Saturazione</label>
          </td>
          <td>
            <input
              id="saturation"
              type="number"
              v-model="configuration.hardware.camera.saturation"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="awbMode">Modo compensazione bianco</label>
          </td>
          <td>
            <select
              id="awbMode"
              v-model="configuration.hardware.camera.awbMode"
            >
              <option v-for="i in awbModes" :key="i" :value="i">{{ i }}</option>
            </select>
          </td>
        </tr>
        <tr>
          <td>
            <label for="num-photo">Numero foto per ciclo</label>
          </td>
          <td>
            <input
              id="num-photo"
              type="number"
              v-model="configuration.hardware.camera.numPhotosPerProcess"
            />
          </td>
        </tr>
        <tr>
          <h2>Server</h2>
        </tr>
        <tr>
          <td>
            <label for="imagePath">Path immagini</label>
          </td>
          <td>
            <input
              id="imagePath"
              type="text"
              v-model="configuration.server.photoDirectory"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="serverPort">Porta server (richiede restart)</label>
          </td>
          <td>
            <input
              id="serverPort"
              type="text"
              v-model="configuration.server.port"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="saveAsWebP">Salva come webp</label>
          </td>
          <td>
            <input
              id="saveAsWebP"
              type="checkbox"
              v-model="configuration.server.saveAsWebP"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="useArducamAufotocusCamera">Usa camera arducam</label>
          </td>
          <td>
            <input
              id="useArducamAufotocusCamera"
              type="checkbox"
              v-model="configuration.server.useArducamAufotocusCamera"
            />
          </td>
        </tr>
        <tr>
          <h2>ECommerce</h2>
        </tr>
        <tr>
          <td>
            <label for="baseUrl">Base url</label>
          </td>
          <td>
            <input
              id="baseUrl"
              type="text"
              v-model="configuration.eCommerce.baseUrl"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="tokensEndpoint">Endpoint token</label>
          </td>
          <td>
            <input
              id="tokensEndpoint"
              type="text"
              v-model="configuration.eCommerce.tokensEndpoint"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="itemsListEndpoint">Endpoint lista articoli</label>
          </td>
          <td>
            <input
              id="itemsListEndpoint"
              type="text"
              v-model="configuration.eCommerce.itemsListEndpoint"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="postImageEndpoint">Endpoint post immagini</label>
          </td>
          <td>
            <input
              id="postImageEndpoint"
              type="text"
              v-model="configuration.eCommerce.postImageEndpoint"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="defaultUsername">Default username</label>
          </td>
          <td>
            <input
              id="defaultUsername"
              type="text"
              v-model="configuration.eCommerce.defaultUsername"
            />
          </td>
        </tr>
        <tr>
          <td>
            <label for="defaultPassword">Default password</label>
          </td>
          <td>
            <input
              id="defaultPassword"
              type="text"
              v-model="configuration.eCommerce.defaultPassword"
            />
          </td>
        </tr>
        <tr>
          <td>
            <button type="button" @click="updateConfiguration">Aggiorna</button>
          </td>
        </tr>
      </table>
    </form>
  </div>
</template>

<script>
export default {
  name: 'ConfigurationView',
  data() {
    return {
      error: '',
      message: '',
      configuration: {
        hardware: {
          motorDegreePerStep: 0,
          waitForStep: 0,
          gearRatio: 0,
          camera: {
            height: 0,
            width: 0,
            brightness: 0,
            sharpness: 0,
            contrast: 0,
            awbMode: 'auto',
            saturation: 0,
          },
        },
        server: {
          photoDirectory: '',
          distributionDirectory: '',
          port: 3333,
          saveAsWebP: false,
          useArducamAufotocusCamera: false,
        },
        eCommerce: {
          baseUrl: '',
          tokensEndpoint: '',
          itemsListEndpopint: '',
          postImageEndpoint: '',
          defaultUsername: '',
          defaultPassword: '',
        },
      },
    }
  },
  methods: {
    updateConfiguration() {
      this.error = ''

      fetch('/api/configuration/', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.configuration),
      })
        .then((a) => {
          if (!a.ok) {
            throw Error('Cannot save config')
          }
          return a.json()
        })
        .then((data) => {
          this.message = data
        })
        .catch((error) => {
          console.log(error)
          this.error = error
        })
    },
  },
  computed: {
    awbModes() {
      return [
        'off',
        'auto',
        'sun',
        'cloud',
        'shade',
        'tungsten',
        'fluorescent',
        'incandescent',
        'flash',
        'horizon',
      ]
    },
  },
  mounted() {
    fetch('/api/configuration')
      .then((data) => {
        return data.json()
      })
      .then((config) => {
        this.configuration = config
      })
  },
}
</script>

<style></style>
