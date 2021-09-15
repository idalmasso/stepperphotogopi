<template>
  <div>
    <h1>Processi</h1>
    <span>{{ message }}</span>
    <table>
      <tr>
        <th>Processo</th>
        <th>Scarica</th>
        <th>Elimina</th>
      </tr>
      <tr v-for="process in processes" :key="process">
        <td>
          <router-link
            :to="{
              name: 'EndedProcessViewer',
              params: { processName: process },
            }"
          >
            {{ process }}</router-link
          >
        </td>
        <td>
          <button @click="downloadProcess(process)">Scarica</button>
        </td>
        <td>
          <button @click="deleteProcess(process)">Elimina</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
export default {
  name: 'EndedProcesses',
  data() {
    return {
      processes: [],
      message: '',
    }
  },
  methods: {
    deleteProcess(process) {
      fetch('/api/processes/' + process, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then((a) => {
          if (a.ok) {
            const idx = this.processes.findIndex((val) => val == process)
            
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
    downloadProcess(process){
      fetch('/api/processes/' + process, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then(async(a) => {
          if(a.ok){
            return a.json()
          } else {
            throw Error(await a.text())
          }
        })
        .then((data) => {
          
          var link = document.createElement("a");
          link.download = process;
          link.href = data['value'];
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
          
        })
        .catch(err=> this.message=err)
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
  },
  mounted() {
    this.getProcesses()
  },
}
</script>

<style scoped>
tr:hover {background-color: #ddd;}
</style>
