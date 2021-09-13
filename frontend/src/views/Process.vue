<template>
  <div class="process">
    <h1>Start Process</h1>
    <ProcessControl @setInterval="motorIntervalSet"/>
  </div>
</template>

<script>
// @ is an alias to /src
import ProcessControl from '@/components/ProcessControl.vue'

export default {
  name: 'Process',
  data(){
    return {
      motorIntervalId: ''
    }
  },
  components: {
    ProcessControl,
  },
  methods:{
    motorIntervalSet(value){
      this.motorIntervalId=value;
    },
  },
  beforeRouteLeave (to, from, next) {
    console.log("ProcessView Clearing statusRequest")
      clearInterval(this.motorIntervalId);
      next();
  },
}
</script>
