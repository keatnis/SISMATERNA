<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :label-position="labelPosition" :mobile-mode="mobileMode">

    <b-step-item step="1" label="Registrar embarazada" :clickable="isStepsClickable">
      <h1 class="title has-text-centered">Registrar embarazada</h1>

      <div class></div>
      <h2 class="subtitle">DATOS PERSONALES</h2>
      <div class="column"></div>

      <b-field label="No. Progreso" :label-position="labelPosition">
        <b-input disabled=»disabled» v-model="detalles.placas" placeholder=""></b-input>
      </b-field>

      <b-field label="CURP">
        <b-input v-model="detalles.descripcion" placeholder="Clave Unica de Registro de Población"></b-input>
      </b-field>

      <b-field label="No. Expediente">
        <b-input v-model="detalles.descripcion" placeholder="No. Expediente clinico"></b-input>
      </b-field>

      <b-field label="Nombre" :label-position="labelPosition">
        <b-input v-model="detalles.placas" placeholder="Nombre completo de paciente"></b-input>
      </b-field>

      <b-field label="Fecha de Nacimiento">
        <b-datepicker placeholder="Fecha de Nacimiento" icon="calendar-today" :locale="locale" editable>
        </b-datepicker>
      </b-field>

      <b-field label="Edad" :label-position="labelPosition">
        <b-input disabled=»disabled» v-model="detalles.placas"></b-input>
      </b-field>


      <b-field label="Número  teléfono">
        <b-input v-model="detalles.descripcion" placeholder="Teléfono fijo"></b-input>
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

      <b-field label="Derechohabiencia">
        <b-select placeholder="Derechohabiencia" expanded>
          <option value="Insabi">INSABI</option>
          <option value="Imss">IMSS</option>
          <option value="Issste">ISSSTE</option>
          <option value="Sedena">SEDENA</option>
          <option value="Pemex">PEMEX</option>
          <option value="Marina">MARINA</option>
          <option value="Otro">OTRO</option>
        </b-select>
      </b-field>

      <div class="control">
        <b-field label="¿Emigro?"></b-field>
        <label class="radio">
          <input type="radio" name="answer">
          Si
        </label>
        <label class="radio">
          <input type="radio" name="answer">
          No
        </label>
        <b-input v-model="detalles.descripcion" placeholder="¿Dónde emigró?"></b-input>
      </div>

      <h2 class="subtitle"></h2>
      <h2 class="subtitle"> DOMICILIO</h2>
      <b-field label="Dirección">
        <b-input v-model="detalles.descripcion" placeholder="Dirección con referencia"></b-input>
      </b-field>
      <b-field label="Localidad">
        <b-input v-model="detalles.descripcion" placeholder="Localidad"></b-input>
      </b-field>

      <b-field label="Municipio">
        <b-input v-model="detalles.descripcion" placeholder="Municipio"></b-input>
      </b-field>

      <h2 class="subtitle"></h2>
      <h2 class="subtitle">DATOS CLÍNICOS</h2>


      <div class="columns">
        <div class="column">
          <b-field label="Detenciones"></b-field>
          <div class="field">
            <div class="b-checkbox">
              <input id="checkbox1" class="styled" checked type="checkbox">
              <label for="checkbox1">
                Exploración clínica de mama normal
              </label>
            </div>
            <div class="b-checkbox">
              <input id="checkbox2" class="styled" checked type="checkbox">
              <label for="checkbox2">
                Exploración clínica de mama alterado
              </label>
              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox">
                <label for="checkbox1">
                  Papanicolau normal
                </label>
              </div>
              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox">
                <label for="checkbox1">
                  Papanicolau con lesiones
                </label>
              </div>
              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox">
                <label for="checkbox1">
                  Papanicolau cancer IN SITU
                </label>
              </div>


              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox">
                <label for="checkbox1">
                  Papanicolau invasor
                </label>
              </div>
              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox">
                <label for="checkbox1">
                  No se realizo
                </label>
              </div>

            </div>

            <div class="column"></div>
            <div class="control">
              <b-field label="¿Asistió a consulta pregestacional?"></b-field>
              <label class="radio">
                <input type="radio" name="answer">
                Si
              </label>
              <label class="radio">
                <input type="radio" name="answer">
                No
              </label>
            </div>
            <b-field label="">
              <b-datepicker placeholder="Fecha de consulta" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
            </b-field>
          </div>
          <div class="columns">

            <div class="column">
              <b-field label="Comorbilidades"></b-field>
              <div class="field">
                <div class="b-checkbox">

                  <input id="checkbox1" class="styled" checked type="checkbox">
                  <label for="checkbox1">
                    SOBRE PESO

                    <input id="checkbox2" class="styled" checked type="checkbox">
                    <label for="checkbox2">
                      OBESIDAD
                    </label>
                  </label>
                  <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  DIABETES MELLITUS
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  HIPERTENSIÓN ARTERIAL
                </label>
                </div>
              
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  CARDIOPATÍA
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  EPILEPSIA
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  SÍFILITIS
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  VIH
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  HEPATITIS
                </label>
                <div>
                </div>
                <div>
                  <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  ARTRITIS REUMATOIDE
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  LUPUS
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  CANCER

                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  TUBERCULOSIS
                </label>
                </div>
              <div>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                COVID-19
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  INFECCIÓN DE VIAS URINARIAS
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  CERVICO VAGINITIS
                </label>
              </div>
              <div>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  ANEMIA
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  OTRO
                </label>
                <input id="checkbox2" class="styled" checked type="checkbox">
                <label for="checkbox2">
                  NINGUNO
                </label>
              </div>
              </div>
            </div>
          </div>
          <b-field label="Gestas">
        <b-input v-model="detalles.descripcion" placeholder="Numero de gestas"></b-input>
      </b-field>

      <b-field label="Paras">
        <b-input v-model="detalles.descripcion" placeholder="Numero de paras"></b-input>
      </b-field>
        </div>
      </div>
      <b-field label="Abortos">
        <b-input v-model="detalles.descripcion" placeholder="Numero de abortos"></b-input>
      </b-field>
      <b-field label="Cesareas">
        <b-input v-model="detalles.descripcion" placeholder="Numero de cesareas"></b-input>
      </b-field>

      <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
        <b-datepicker placeholder="Fecha de Nacimiento" icon="calendar-today" :locale="locale" editable>
        </b-datepicker>
      </b-field>

      

    </b-step-item>

    <b-step-item step="2" label="Registrar embarazada" :clickable="isStepsClickable"
      :type="{ 'is-success': isProfileSuccess }">
      <h1 class="title has-text-centered">Registrar embarazada</h1>
      <div class="column"> </div>
      <div class="control">
        <b-field label="¿Emigro?"></b-field>
        <label class="radio">
          <input type="radio" name="answer">
          Si
        </label>
        <label class="radio">
          <input type="radio" name="answer">
          No
        </label>
        <b-input v-model="detalles.descripcion" placeholder="¿Dónde emigró?"></b-input>
      </div>
    </b-step-item>

    <b-step-item step="3" :visible="showSocial" label="Registrar embarazada" :clickable="isStepsClickable">
      <h1 class="title has-text-centered">Registrar embarazada</h1>
      Lorem ipsum dolor sit amet.
    </b-step-item>

    <b-step-item :step="showSocial ? '4' : '4'" label="Trigge" :clickable="isStepsClickable" disabled>
      <h1 class="title has-text-centered">Trigge</h1>
      Lorem ipsum dolor sit amet.
      <b-field>
        <b-button @click="guardar()" type="is-success">Guardar</b-button>
        <router-link :to="{ name: 'Vehiculos' }" class="button is-info ml-2">
          Volver
        </router-link>
      </b-field>
    </b-step-item>

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