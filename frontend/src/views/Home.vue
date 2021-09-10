<template>
  <div class="home">
    <MotorControl />
  </div>
</template>

<script>
// @ is an alias to /src
import MotorControl from '@/components/MotorControl.vue'

export default {
  name: 'Home',
  components: {
    MotorControl,
  },
}
</script>
