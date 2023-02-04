<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :label-position="labelPosition" :mobile-mode="mobileMode">
    <div class="has-text-centered">
      <article class="panel is-info centered">
        <div class="panel-heading">
          Registrar puerpera
        </div>
      </article>
    </div>
    <div>
      <div class="columns">
        <div class="column is-half">

          <b-field label="No. Progreso" :label-position="labelPosition">
            <b-input disabled="»disabled»" v-model="detalles.id" placeholder=""></b-input>
          </b-field>
          <b-field label="Nombre de puerpera" :label-position="labelPosition">
            <b-input v-model="detalles.NombrePuerpera" placeholder="Nombre completo"></b-input>
          </b-field>

          <b-field label=" Consulta de 7 días">
            <b-datepicker placeholder="Fecha consulta de 7 días" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>



          <b-field label="Consulta de 40 días">
            <b-datepicker placeholder="Fecha consulta de 40 días" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>

          <b-field label="Atención de parto">
            <b-select placeholder="¿Quién atendió el parto?" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Resolución del embarazo">
            <b-select placeholder="Resolución del embarazo" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="APEO">
            <b-select placeholder="APEO" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>
          <b-field label="Consulta pregestacional"></b-field>
          <label class="radio">
            <input type="radio" name="answer" />
            Si
          </label>
          <label class="radio">
            <input type="radio" name="answer" />
            No
          </label>
          <b-field label="Derechohabiencia">
            <b-select placeholder="Derechohabiencia" disabled expanded>
              <option value="Insabi">INSABI</option>
              <option value="Imss">IMSS</option>
              <option value="Issste">ISSSTE</option>
              <option value="Sedena">SEDENA</option>
              <option value="Pemex">PEMEX</option>
              <option value="Marina">MARINA</option>
              <option value="Otro">Otro</option>
            </b-select>
          </b-field>
        </div>
        <div class="column is-half">
          <b-field label="CURP">
            <b-input v-model="detalles.descripcion" placeholder="Clave Unica de Registro de Población" icon="magnify"></b-input>
          </b-field>
          <b-field label="Edad">
            <b-input disabled="»disabled»" v-model="detalles.Edad"></b-input>
          </b-field>

          <b-field label="Consulta de 28 días">
            <b-datepicker placeholder="Fecha consulta de 28 días" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>
          <b-field label="Signos de alarma">
            <b-select placeholder="Signos" expanded>
              <option value="Fiebre">Fiebre</option>
              <option value="DolorPelvico">Dolor Pélvico</option>
              <option value="FlujoVaginal">Flujo Vaginal</option>
              <option value="FiebreDolor">Fiebre/Dolor Pelvico</option>
              <option value="FiebreFlujo">Fiebre/Flujo Vaginal</option>
              <option value=">DolorFlujo">Dolor Pelvico/Flujo Vaginal</option>
              <option value="FiebreDolorFlujo">Fiebre/Dolor Pelvico/Flujo Vaginal</option>
              <option value="Cefalea">Cefalea</option>
              <option value="NingunoSigno">Ninguno</option>
            </b-select>
          </b-field>

          <b-field label="Lugar del parto">
            <b-select placeholder="¿Dónde atendió el parto?" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>



          <b-field label="No. de producto">
            <b-select placeholder="Bo. de producto" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>



          <b-field label="Puerpera aceptante">
            <b-select placeholder="Puerpera acentante" expanded>
              <option value="Nahual">Nahual</option>
              <option value="Mixteco">Tu´un Savi (Mixteco)</option>
              <option value="Mephaa">Meephaa (Tlapaneco)</option>
              <option value="Amuzgo">Amuzgo</option>
              <option value="Ninguno">Ninguno</option>
            </b-select>
          </b-field>


          <b-field label="Fecha de parto">
            <b-datepicker placeholder="Fecha del evento" icon="calendar-today" :locale="locale" editable>
            </b-datepicker>
          </b-field>
          <b-field>
            <b-switch v-model="usarFechaYHoraActual">Usar fecha y hora actual ({{ fechaYHoraActual }})</b-switch>
          </b-field>
          <b-field>
            <b-button @click="guardar()" type="is-success">Guardar</b-button>

          </b-field>
        </div>

        <b-field label="Selecciona manualmente" v-show="!usarFechaYHoraActual">
          <b-datetimepicker v-model="detalles.fechaEntrada" rounded placeholder="Clic aquí para seleccionar"
            icon="calendar-today" locale="es-MX" :datetime-formatter="formatearFecha"
            :timepicker="{ enableSeconds: false, hourFormat: '24' }" inline>
            <template #left>
              <b-button label="Ahora" type="is-primary" icon-left="clock" @click="detalles.fechaEntrada = new Date()" />

            </template>

            <template #right>
              <b-button label="Limpiar" type="is-danger" icon-left="close" outlined
                @click="detalles.fechaEntrada = null" />
            </template>
          </b-datetimepicker>
        </b-field>

      </div>
    </div>


    <template v-if="customNavigation" #navigation="{ previous, next }">
      <b-button outlined type="is-danger" icon-pack="fas" icon-left="backward" :disabled="previous.disabled"
        @click.prevent="previous.action">
        Previous
      </b-button>
      <b-button outlined type="is-success" icon-pack="fas" icon-right="forward" :disabled="next.disabled"
        @click.prevent="next.action">
        Next
      </b-button>
    </template>
  </b-steps>


</template>
<script>
//labelPosition: 'on-border',
import Utiles from "../services/Utiles";
import DialogosService from "../services/DialogosService";
import EmbarazadaService from "../services/EmbarazadaService";
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
      const respuesta = await EmbarazadaService.agregarVehiculo(cargaUtil);
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