<template>
  <!-- <div class="columns"> -->
    <b-steps
      v-model="activeStep"
      :animated="isAnimated"
      :rounded="isRounded"
      :has-navigation="hasNavigation"
      :icon-prev="prevIcon"
      :icon-next="nextIcon"
      :label-position="labelPosition"
      :mobile-mode="mobileMode"
    >

      <b-step-item step="1" label="Datos Generales de la embarazada" :clickable="isStepsClickable">
        <h1 class="title has-text-centered">Datos Generales de la embarazada</h1>

        <div class="column">
          
          <b-field label="Nombre" :label-position="labelPosition">
            <b-input
              v-model="detalles.placas"
              placeholder="Nombre de Embarazada"
            ></b-input>
          </b-field>
          <b-field label="CURP">
            <b-input
              v-model="detalles.descripcion"
              placeholder="Clave Unica de Registro de Población"
            ></b-input>
          </b-field>

          <b-field label="Fecha de Nacimiento">
            <b-datepicker
              placeholder="Fecha de Nacimiento"
              icon="calendar-today"
              :locale="locale"
              editable
            >
            </b-datepicker>
          </b-field>

          <b-field label="Número telefónico">
            <b-input placeholder="Número " type="number" min="10" max="12">
            </b-input>
          </b-field>
          <b-field label="Lengua Indígena">
            <b-select placeholder="Lengua Indígena" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>
          <h2 class="subtitle"></h2>
          <h2 class="subtitle">--- DOMICILIO ---</h2>
          <b-field label="Dirección">
            <b-input
              v-model="detalles.descripcion"
              placeholder="Dirección"
            ></b-input>
          </b-field>
          <b-field label="Localidad">
            <b-input
              v-model="detalles.descripcion"
              placeholder="Localidad"
            ></b-input>
          </b-field>

          <b-field label="Municipio">
            <b-input
              v-model="detalles.descripcion"
              placeholder="Municipio"
            ></b-input>
          </b-field>

          <b-field>
            <b-switch v-model="usarFechaYHoraActual"
              >Usar fecha y hora actual ({{ fechaYHoraActual }})</b-switch
            >
          </b-field>

          <b-field
            label="Selecciona manualmente"
            v-show="!usarFechaYHoraActual"
          >
            <b-datetimepicker
              v-model="detalles.fechaEntrada"
              rounded
              placeholder="Clic aquí para seleccionar"
              icon="calendar-today"
              locale="es-MX"
              :datetime-formatter="formatearFecha"
              :timepicker="{ enableSeconds: false, hourFormat: '24' }"
              inline
            >
              <template #left>
                <b-button
                  label="Ahora"
                  type="is-primary"
                  icon-left="clock"
                  @click="detalles.fechaEntrada = new Date()"
                />
              </template>

              <template #right>
                <b-button
                  label="Limpiar"
                  type="is-danger"
                  icon-left="close"
                  outlined
                  @click="detalles.fechaEntrada = null"
                />
              </template>
            </b-datetimepicker>
          </b-field>
          <b-field>
            <b-button @click="guardar()" type="is-success">Guardar</b-button>
            <router-link
              :to="{ name: 'Vehiculos' }"
              class="button is-info ml-2"
            >
              Volver
            </router-link>
          </b-field>
        </div>
      </b-step-item>

      <b-step-item
        step="2"
        label="Datos generales de la embarazada"
        :clickable="isStepsClickable"
        :type="{ 'is-success': isProfileSuccess }"
      >
        <h1 class="title has-text-centered">Datos generales de la embarazada</h1>
        Lorem ipsum dolor sit amet.
      </b-step-item>

      <b-step-item
        step="3"
        :visible="showSocial"
        label="Datos perosnales de la embarazada"
        :clickable="isStepsClickable"
      >
        <h1 class="title has-text-centered">Da</h1>
        Lorem ipsum dolor sit amet.
      </b-step-item>

      <b-step-item
        :step="showSocial ? '4' : '4'"
        label="Trigge"
        :clickable="isStepsClickable"
        disabled
      >
        <h1 class="title has-text-centered">Trigge</h1>
        Lorem ipsum dolor sit amet.
      </b-step-item>

      <template v-if="customNavigation" #navigation="{ previous, next }">
        <b-button
          outlined
          type="is-danger"
          icon-pack="fas"
          icon-left="backward"
          :disabled="previous.disabled"
          @click.prevent="previous.action"
        >
          Previous
        </b-button>
        <b-button
          outlined
          type="is-success"
          icon-pack="fas"
          icon-right="forward"
          :disabled="next.disabled"
          @click.prevent="next.action"
        >
          Next
        </b-button>
      </template>
    </b-steps>

   
</template>
<script>
//labelPosition: 'on-border',
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import VehiculosService from "../services/VehiculosService";
export default {
  data: () => ({
    selected: null,
    fechaYHoraActual: null,
    usarFechaYHoraActual: true,
    detalles: {
      placas: "",
      descripcion: "",
      fechaEntrada: null,
    },
  }),
  computed: {
    selectedString() {
      return this.selected ? this.selected.toDateString() : "";
    },
  },
  mounted() {
    this.detalles.fechaEntrada = new Date();
    this.refrescarFechaYHoraActual();
    setInterval(this.refrescarFechaYHoraActual, 500);
  },
  methods: {
    refrescarFechaYHoraActual() {
      this.fechaYHoraActual = Utiles.obtenerFechaYHoraActual();
    },
    formatearFecha(fecha) {
      return Utiles.obtenerFechaYHora(fecha);
    },
    async guardar() {
      if (!this.usarFechaYHoraActual && !this.detalles.fechaEntrada) {
        return DialogosService.mostrarNotificacionError(
          "Selecciona una fecha y hora"
        );
      }
      const cargaUtil = {
        placas: this.detalles.placas,
        descripcion: this.detalles.descripcion,
        fechaEntrada: Utiles.obtenerFechaYHora(this.detalles.fechaEntrada, "T"),
      };
      if (this.usarFechaYHoraActual) {
        cargaUtil.fechaEntrada = Utiles.obtenerFechaYHoraActual("T");
      }
      const respuesta = await VehiculosService.agregarVehiculo(cargaUtil);
      if (respuesta) {
        DialogosService.mostrarNotificacionExito("Vehículo registrado");
        this.detalles = {
          placas: "",
          descripcion: "",
          fechaEntrada: new Date(),
        };
      }
    },
  },
};
</script>