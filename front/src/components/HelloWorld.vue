<template>
  <div class="hello">
    <label for="band">Nombre de la banda</label><br>
    <input type="text" id="band" name="bandName" v-model="band"><br>
    <br>
    <button type="button" class="btn btn-primary" v-on:click='getTracks(this.band)' :disabled="isSpinner">Buscar</button>
    <br>
    <br>
    <div class="spinner-border" role="status" v-if="isSpinner">
      <span class="sr-only">Loading...</span>
    </div>
    <TableBand v-if="isActive" v-bind:songs="this.songs"></TableBand>
  </div>
</template>

<script>
import axios from "axios";
import TableBand from "./TableBand.vue";
export default {
  name: 'HelloWorld',

  data() {
    return {
      isSpinner: false,
      isActive: false,
      songs: null
    }
  },
  components: {
    TableBand,
  },
  props: {
  },
  methods: {
    getTracks: async function (band) {
      this.isActive = false;
      this.isSpinner = true;
      axios
        .get('http://localhost:3000/search_tracks?name=' + encodeURIComponent(band))
        .then(response => {
          this.isActive = true;
          this.songs = response.data.canciones;
          this.isSpinner = false;
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
