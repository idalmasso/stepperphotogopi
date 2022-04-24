<template>
  <div>
    <ECommerceLoginDialog
      v-if="showLogin"
      @close="showLogin = false"
      @logged="loggedInOk"
      :authUrl="authUrl"
      :defaultUsername="configuration.eCommerce.defaultUsername"
      :defaultPassword="configuration.eCommerce.defaultPassword"
    >
    </ECommerceLoginDialog>
    <ItemSelectorDialog
      v-if="showItemSelector"
      @close="showItemSelector = false"
      @selected="selectedItem"
      :itemsUrl="itemsUrl"
      :postImageUrl="postImageUrl"
    >
    </ItemSelectorDialog>
    <h1>Processi</h1>
    <span>{{ message }}</span>
    <span v-if="numErrors != 0 || numOk != 0"
      >Upload errors: {{ numErrors }}, Ok: {{ numOk }}</span
    >
    <table>
      <tr>
        <th>Processo</th>
        <th>Scarica</th>
        <th>Elimina</th>
        <th>Inserisci in ECOMMERCE</th>
      </tr>
      <tr v-for="process in processes" :key="process.name">
        <td>
          <router-link
            :to="{
              name: 'EndedProcessViewer',
              params: {
                processName: process.name,
                numImages: process.numFiles,
              },
            }"
          >
            {{ process.name }}</router-link
          >
        </td>
        <td>
          <button @click="downloadProcess(process)">Scarica</button>
        </td>
        <td>
          <button @click="deleteProcess(process)">Elimina</button>
        </td>
        <td>
          <button @click="uploadProcess(process)">Inserisci nel sito</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
import ECommerceLoginDialog from '@/components/ECommerceLoginDialog.vue'
import ItemSelectorDialog from '@/components/ItemSelectorDialog.vue'
export default {
  name: 'EndedProcesses',
  components: { ECommerceLoginDialog, ItemSelectorDialog },
  data() {
    return {
      numErrors: 0,
      numOk: 0,
      processes: [],
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
          },
        },
        server: {
          photoDirectory: '',
          distributionDirectory: '',
          port: 3333,
        },
        eCommerce: {
          baseUrl: '',
          tokensEndpoint: '',
          itemsListEndpopint: '',
          postImageEndpoint: '',
        },
      },
      showLogin: false,
      sendingProcess: null,
      sendingItem: '',
      showItemSelector: false,
      token: '',
    }
  },
  methods: {
    deleteProcess(process) {
      fetch('/api/processes/' + process.name, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then((a) => {
          if (a.ok) {
            const idx = this.processes.findIndex((val) => val.name == process)

            if (idx > 0) {
              this.processes.splice(idx, 1)
            }
          }
          return a.json()
        })
        .then((data) => {
          this.message = data
          this.getProcesses()
        })
    },
    downloadProcess(process) {
      fetch('/api/processes/' + process.name, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then(async (a) => {
          if (a.ok) {
            return a.json()
          } else {
            throw Error(await a.text())
          }
        })
        .then((data) => {
          var link = document.createElement('a')
          link.download = process
          link.href = data['value']
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link)
        })
        .catch((err) => (this.message = err))
    },
    getProcesses() {
      fetch('/api/processes')
        .then((a) => {
          return a.json()
        })
        .then((data) => {
          this.processes = data['value']
        })
    },
    uploadProcess(process) {
      this.numErrors = 0
      this.numOk = 0
      this.showLogin = true
      this.sendingProcess = process
    },
    loggedInOk(token) {
      console.log(token)
      this.token = token
      this.showLogin = false
      this.showItemSelector = true
    },
    selectedItem(item) {
      console.log(item)
      this.showItemSelector = false
      this.sendingItem = item

      const upload = (i) => {
        console.log('Starting fetch' + i)
        var is_webp = false
        return fetch(
          '/process-images/' + this.sendingProcess.name + '/' + i + '.webp'
        )
          .then((response) => {
            if (response.ok) {
              is_webp = true
              return response.blob()
            } else {
              return fetch(
                '/process-images/' + this.sendingProcess.name + '/' + i + '.jpg'
              )
              .then((response) => {
                if (response.ok) {
                  is_webp = false
                  return response.blob()
                } else {
                  throw 'Cannot get image ' + i
                }
              }
            }
          })
          .then((bData) => {
            let reader = new FileReader()
            reader.readAsDataURL(bData)
            reader.onload = () => {
              const src = reader.result
              fetch(this.postImagesUrl + '/' + this.sendingItem + '/' + i, {
                method: 'POST',
                headers: {
                  'Content-Type': 'application/json',
                  Authorization: 'Bearer ' + this.token,
                },
                body: JSON.stringify({ image: src, is_webp: is_webp }), // This is your file object
              })
                .then(
                  (response) => response.json() // if the response is a JSON object
                )
                .then(() => {
                  this.numOk++
                })
            }
          })
          .catch(() => {
            this.numErrors++
          })
      }
      var fetches = []
      for (var i = 1; i <= this.sendingProcess.numFiles; i++) {
        fetches.push(upload(i))
      }
    },
  },
  computed: {
    authUrl() {
      return (
        this.configuration.eCommerce.baseUrl +
        this.configuration.eCommerce.tokensEndpoint
      )
    },
    itemsUrl() {
      return (
        this.configuration.eCommerce.baseUrl +
        this.configuration.eCommerce.itemsListEndpoint
      )
    },
    postImagesUrl() {
      return (
        this.configuration.eCommerce.baseUrl +
        this.configuration.eCommerce.postImageEndpoint
      )
    },
  },
  mounted() {
    this.getProcesses()
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

<style scoped>
tr:hover {
  background-color: #ddd;
}
</style>
