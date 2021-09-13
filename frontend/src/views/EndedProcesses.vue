<template>
  <div>
    <h1>Processes</h1>
    <li v-for="process in processes" :key="process">
      <router-link :to="{name:'EndedProcessViewer' , params:{processName: process}}"  > {{ process }}</router-link>
    </li>
  </div>
</template>

<script>
export default {
  name: 'EndedProcesses',
  data() {
    return {
      processes: [],
    }
  },
  mounted() {
    fetch('/api/processes')
      .then((a) => {
        return a.json()
      })
      .then((data) => {
        console.log(data)
        this.processes = data['value']
      })
  },
}
</script>

<style></style>
