<template>
  <div>
    <table class="table">
      <thead>
        <tr>
          <th scope="col">id</th>
          <th scope="col">album</th>
          <th scope="col">tema</th>
          <th scope="col">preview</th>
          <th scope="col">fecha lanzamiento</th>
          <th scope="col">valor(USD)</th>
        </tr>
        <tr v-for="(song, index) in this.songs" :key="index">
            <td>
              {{ song.cancion_id}}
            </td>
            <td>
              {{ song.nombre_album}}
            </td>
            <td>
              {{ song.nombre_tema}}
            </td>
            <td>
              {{song.preview_url}}
            </td>
            <td>
              {{song.fecha_lanzamiento}}
            </td>
            <td>
              {{song.precio.valor}}
            </td>
            

        </tr>
      </thead>
      <tbody>
      </tbody>
    </table>
  </div>
</template>
  
<script>
export default {
  name: 'TableBand',
  props: {
    songs: Array
  },
  methods: {
  },
  mounted() {
    this.tableData=this.songs
  },
  data() {
    return {
      tableData: this.songs
    }
  }
}

</script>
  
<style>

</style>
  