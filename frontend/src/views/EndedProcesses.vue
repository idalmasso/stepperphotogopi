<template>
  <div>
    <h1>Processes</h1>
    <li v-for="process in processes" :key="process">
      <a :href="processImagesUrl(process)">{{ process }}</a>
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
  methods: {
    processImagesUrl(process) {
      return '/process-images/' + process
    },
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
