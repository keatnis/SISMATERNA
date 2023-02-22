<template>
  <section>
  
   <b-field label="Municipios"   >
            <b-select placeholder="Seleccionar"
                v-model="selectedOptions" :onChange="obtenerLocalidades()">
                <option
                    v-for="option in municipios"
                    :value="{ id: option.id_municipio, text: option.nombreMunicipio }"
                    :key="option.id">
                    {{ option.nombreMunicipio }}

                                    
                </option>
                
            </b-select>
        </b-field>
        <b-field label="Localidades"   >
            <b-select placeholder="Seleccionar localidad"
                v-model="selectedOptions2">
                <option
                    v-for="option in localidades"
                    :value="{ id: option.id_localidad, text: option.NombreLocalidad}"
                    :key="option.id"
                  >
                    {{ option.NombreLocalidad}}
                  
                                    
                </option>
                
            </b-select>

        </b-field>
        <b-field>
       
       
      </b-field>
        <p class="content"><b>selected</b>: {{ selectedOptions.id}}</p>
    </section>
</template>

<script>

import DialogosService from "../services/DialogosService";
import EmbarazadaService from "../services/EmbarazadaService";
export default {
  data() {
    return {
 
      municipios:[],
      localidades:[],
      selectedOptions: [],
      selectedOptions2:[],
        columns: [
      {
          field: "id",
          label: "No. Progreso",
          width: "100",
          numeric: true,
          searchable: true,
          centered: true,
        },
        {
          field: "noExpediente",
          label: "Nombre completo",
          searchable: true,
          centered: true,
        },
        {
          field: "nombre",
          label: "SDG",
          centered: true,
        },
        {
          field: "curp",
          label: "Fecha de consutal",
          searchable: true,
          centered: true,
        },
        {
          field: "telefono",
          label: "telefono",
          searchable: true,
          centered: true,
        },
    ]
  }
},

  async mounted() {
    await this.obtenerMunicipio();
   
   
  
  },
  methods: {
    async obtenerMunicipio() {
     
      try {
        if (typeof this.municipioSeleccionado.id === "undefined") {
          return;
        } else {
          this.localidades = await EmbarazadaService.obteneLocalidadesById(
            this.municipioSeleccionado.id
          );
        }
      } catch (e) {
       DialogosService.mostrarNotificacionError(
         "No se pudo obtener la lista de municipios..."
       );
     } 
     
   },
   async obtenerLocalidades(){
    try{
      this.localidades = await EmbarazadaService.obteneLocalidadesById(this.selectedOptions.id);

    }catch(e){
      DialogosService.mostrarNotificacionError("error localidades")
    }
   }
   
  },
  
};
</script>
<style>
#tabla {
  justify-content: center;
}
</style> 