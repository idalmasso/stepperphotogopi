<template>
  <div>
    <h1>Processi</h1>
    <span>{{ message }}</span>
    <table>
      <tr>
        <th>Processo</th>
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
          <button @click="deleteProcess(process)">DELETE</button>
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
            console.log(idx)
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
    getProcesses() {
      fetch('/api/processes')
        .then((a) => {
          console.log(a)
          return a.json()
        })
        .then((data) => {
          console.log(data)
          this.processes = data['value']
        })
    },
  },
  mounted() {
    this.getProcesses()
  },
}
</script>

<style></style>
