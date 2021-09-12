<template>
  <div class="process">
    <ProcessControl />
  </div>
</template>

<script>
// @ is an alias to /src
import ProcessControl from '@/components/ProcessControl.vue'

export default {
  name: 'Process',
  components: {
    ProcessControl,
  },
}
</script>
