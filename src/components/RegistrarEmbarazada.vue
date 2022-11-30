<template>
  <!-- <div class="columns"> -->
  <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
    :icon-prev="prevIcon" :icon-next="nextIcon" :label-position="labelPosition" :mobile-mode="mobileMode">

    <b-step-item step="" label="" :clickable="isStepsClickable">
      <div class="column has-text-centered">
        <article class="panel is-info centered">
          <p class="panel-heading">
            Registrar embarazada
          </p>
        </article>

      </div>
      <div>
        <div class="columns">
          <div class="column is-one-third">

            <b-field label="No. Progreso" :label-position="labelPosition">
              <b-input disabled="»disabled»" v-model="detalles.placas" placeholder=""></b-input>
            </b-field>

            <b-field label="Nombre" :label-position="labelPosition">
              <b-input v-model="detalles.placas" placeholder="Nombre completo de paciente"></b-input>
            </b-field>
            <b-field label="Número  teléfono">
              <b-input v-model="detalles.descripcion" placeholder="Teléfono fijo"></b-input>
            </b-field>
            <b-field label="Municipio">
              <b-input v-model="detalles.descripcion" placeholder="Municipio"></b-input>
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

            <b-field label="Comorbilidades">
              <section>
                <div class="block">
                  <b-checkbox v-model="checkboxGroup" native-value="sobrepeso">
                    Sobre peso
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Obesidad">
                    Obesidad
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Diabetes Mellitus">
                    Diabetes Mellitus
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Hipertensión arterial">
                    Hipertensión arterial
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Cardiopatía">
                    Cardiopatía
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Epilepsia ">
                    Epilepsia
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Sifilis ">
                    Sifilis
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="VIH ">
                    VIH
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Hepatitis">
                    Hepatitis
                  </b-checkbox>
                  <b-checkbox v-model="checkboxGroup" native-value="Ninguno">
                    Ninguno
                  </b-checkbox>
                </div>
                <p class="content">
                  <b>Selección:</b>
                  {{ checkboxGroup }}
                </p>
              </section>
            </b-field>
            <b-field label="Gestas">
              <b-input v-model="detalles.descripcion" placeholder="Numero de gestas"></b-input>
            </b-field>
            <b-field label="Cesareas">
              <b-input v-model="detalles.descripcion" placeholder="Numero de cesareas"></b-input>
            </b-field>
          </div>

          <div class="column is-one-third">
            <b-field label="CURP">
              <b-input v-model="detalles.descripcion" placeholder="Clave Unica de Registro de Población"></b-input>
            </b-field>
            <b-field label="Fecha de nacimiento">
              <b-datepicker placeholder="Fecha de nacimiento" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
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
            <b-field label="Localidad">
              <b-input v-model="detalles.descripcion" placeholder="Localidad"></b-input>
            </b-field>

            <b-field label="Detenciones"></b-field>
            <div class="field">
              <div class="b-checkbox">
                <input id="checkbox1" class="styled" checked type="checkbox" />
                <label for="checkbox1">
                  Exploración clínica de mama normal
                </label>
              </div>
              <div class="b-checkbox">
                <input id="checkbox2" class="styled" checked type="checkbox" />
                <label for="checkbox2">
                  Exploración clínica de mama alterado
                </label>
                <div class="b-checkbox">
                  <input id="checkbox1" class="styled" checked type="checkbox" />
                  <label for="checkbox1"> Papanicolau normal </label>
                </div>
                <div class="b-checkbox">
                  <input id="checkbox1" class="styled" checked type="checkbox" />
                  <label for="checkbox1"> Papanicolau con lesiones </label>
                </div>
                <div class="b-checkbox">
                  <input id="checkbox1" class="styled" checked type="checkbox" />
                  <label for="checkbox1"> Papanicolau cancer IN SITU </label>
                </div>

                <div class="b-checkbox">
                  <input id="checkbox1" class="styled" checked type="checkbox" />
                  <label for="checkbox1"> Papanicolau invasor </label>
                </div>
                <div class="b-checkbox">
                  <input id="checkbox1" class="styled" checked type="checkbox" />
                  <label for="checkbox1"> No se realizo </label>
                </div>
              </div>

              <b-field label="Paras">
                <b-input v-model="detalles.descripcion" placeholder="Numero de paras"></b-input>
              </b-field>
              <b-field label="Abortos">
                <b-input v-model="detalles.descripcion" placeholder="Numero de abortos"></b-input>
              </b-field>
            </div>

          </div>
          <div class="column is-one-fifthr">
            <b-field label="No. Expediente clínico">
              <b-input v-model="detalles.descripcion" placeholder="No. Expediente clínico"></b-input>
            </b-field>
            <b-field label="Edad" :label-position="labelPosition">
              <b-input disabled="»disabled»" v-model="detalles.placas"></b-input>
            </b-field>
            <div class="control">
              <b-field label="¿Emigro?"></b-field>
              <label class="radio">
                <input type="radio" name="answer" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="answer" />
                No
              </label>
              <b-input v-model="detalles.descripcion" placeholder="¿Dónde emigró?"></b-input>
            </div>
            <b-field label="Dirección">
              <b-input v-model="detalles.descripcion" placeholder="Dirección con referencia"></b-input>
            </b-field>
            <div class="control">
              <b-field label="¿Asistió a consulta pregestacional?"></b-field>
              <label class="radio">
                <input type="radio" name="answer" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="answer" />
                No
              </label>
            </div>
            <b-field label="">
              <b-datepicker placeholder="Fecha de consulta" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
            </b-field>

            <b-field label="Fecha de ultimo evento( Parto, cesarea, aborto)">
              <b-datepicker placeholder="Fecha de ultimo evento" icon="calendar-today" :locale="locale" editable>
              </b-datepicker>
            </b-field>

            <div class="control">
              <b-field label="¿Presentó complicaciones en el embarazo previo?"></b-field>
              <label class="radio">
                <input type="radio" name="answer" />
                Si
              </label>
              <label class="radio">
                <input type="radio" name="answer" />
                No
              </label>
              <b-select placeholder="¿Qué complicaciones presento?" expanded>
                <option value="Hemorragia ">Hemorragia </option>
                <option value="Retención de placenta">Retención de placenta</option>
                <option value="Placenta previa">Placenta previa</option>
                <option value="Diabetes gestacional">Diabetes gestacional</option>
                <option value="Preeclampsia">Preeclampsia</option>
                <option value="Preeclampsia">Eclampsia </option>
                <option value="Preeclampsia">Covid + embarazo</option>
                <option value="Preeclampsia">Otra</option>
                <option value="Preeclampsia">Ninguno</option>
              </b-select>
            </div>

          </div>
        </div>
      </div>
    </b-step-item>


    <b-step-item step="" label="" :clickable="isStepsClickable" :type="{ 'is-success': isProfileSuccess }">
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha de ultima menstruación">
                <b-datepicker placeholder="Fecha de ultima menstruación" icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de consutal">
                <b-datepicker placeholder="Fecha de consulta " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
              <div class="control">
                <b-field label="Rubeola"></b-field>
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  Si se aplicó previo al embarazo
                </label>
                <label class="radio">
                  <input type="radio" name="Rubeola" />
                  No se cuenta con vacuna
                </label>
              </div>
              <div class="control">
                <b-field label="Fecha de vacunación TD refuerzo">
                  <b-datepicker placeholder="Fecha de vacunación TD refuerzo" icon="calendar-today" :locale="locale"
                    editable>
                  </b-datepicker>

                </b-field>

              </div>
              <b-field label="Vacuna COVID-19">
                <label class="radio">
                  <input type="radio" name="answer" />
                  Se aplico
                </label>
                <label class="radio">
                  <input type="radio" name="answer" />
                  No se aplico
                </label>
              </b-field>
              <b-field label="BIOMETRIA EMATICA">
                <b-select placeholder="BIOMETRIA" expanded>
                  <option value="HB12">HB >=12</option>
                  <option value="HBMAY9">HB >=9</option>
                  <option value="HBMEN9">HB MENOR A 9 </option>
                  <option value="Ninguno">NINGUNO</option>
                </b-select>
              </b-field>

              <b-field label="VIH">
                <b-select placeholder="VIH" expanded>
                  <option value="ReactivoVIH">REACTIVO</option>
                  <option value="NoReactivoVIH">NO REACTIVO</option>
                  <option value="NoRealizoVIH">NO SE REALIZO</option>
                </b-select>
              </b-field>
              <b-field label="CARACTERÍSTICAS FETALES ">
                <b-select placeholder="CARACTERÍSTICAS" expanded>
                  <option value="UnicoVivo">UNICO VIVO</option>
                  <option value="UnicObito">UNICO OBITO</option>
                  <option value="Molar">MOLAR</option>
                  <option value="Monocorial">MONOCORIAL</option>
                  <option value="Bicorionicos">BICORIONICOS</option>
                  <option value="Gemelar1Obitado">GMELAR 1 OBITADO</option>
                </b-select>
              </b-field>
            </div>
          </div>

          <div class="column is-one-third">
            <div class="control">
              <b-field label="Fecha probable de parto">
                <b-datepicker placeholder="Fecha probable de parto " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="No. Total de consulta durante el embarazo" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas"
                placeholder="No. Total de consulta durante el embarazo"></b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD primera">
                <b-datepicker placeholder="Fecha de vacunación TD primera " icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación TDPA">
                <b-datepicker placeholder="Fecha de vacunación TDPA " icon="calendar-today" :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>

            <b-field label="Grupo RH">
              <b-select placeholder="Grupo RH" expanded>
                <option value="A+">A+</option>
                <option value="A-">A-</option>
                <option value="B+">B+</option>
                <option value="B-">B-</option>
                <option value="O+">O+</option>
                <option value="O-">O-</option>
                <option value="AB+">AB+</option>
                <option value="AB-">AB-</option>
                <option value="NingunoRH">NINGUNO</option>
              </b-select>
            </b-field>
            <b-field label="LEUCOCITOS">
              <b-select placeholder="LEUCOCITOS" expanded>
                <option value="Normales">NORMALES</option>
                <option value="DisminuidosLinfocitos">DISMINUIDOS LINFOCITOS</option>
                <option value="AumentadosNeutrofilos">AUMENTADOS NEUTROFILOS</option>
                <option value="AumentadosLinfocitos">AUMENTADOS LOS LINFOCITOS</option>
                <option value="NingunoLEUCO">NINGUNO</option>
              </b-select>
            </b-field>

            <b-field label="ESTADO TIRAS GLUCOSA">
              <b-select placeholder="ESTADO EN EL QUE SE REALIZO" expanded>
                <option value="Ayuno">AYUNO</option>
                <option value="Posprandial1">POSPRANDIAL 1 HORA</option>
                <option value="Posprandial2">POSPRANDIAL 2 HORAS</option>
                <option value="NingunoLEUCO">NINGUNO</option>
              </b-select>
            </b-field>

            <b-field label=" MALFORMACIONES CONGENITAS">
              <b-select placeholder="MALFORMACIONES" expanded>
                <option value="EspinaBifida">ESPINA BIFIDA</option>
                <option value="Meningocele">MENINGOCELE</option>
                <option value="Mielomeningocele">MIELOMENINGOCELE</option>
                <option value="Anencefalia">ANENCEFALIA</option>
                <option value="OtrasMalform">OTRAS</option>
                <option value="NingunaMalform">NINGUNA</option>
              </b-select>
            </b-field>

          </div>
          <div class="column is-one-third">

            <b-field label="Semana de gestación actual" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas" placeholder="Semanan de gestación actual">
              </b-input>
            </b-field>
            <b-field label="No. Total de consulta en el mes" :label-position="labelPosition">
              <b-input disabled=»disabled» v-model="detalles.placas" placeholder="No. Total de consulta en el mes ">
              </b-input>
            </b-field>
            <div class="control">
              <b-field label="Fecha de vacunación TD segunda">
                <b-datepicker placeholder="Fecha de vacunación TD segunda " icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
            <div class="control">
              <b-field label="Fecha de vacunación Influenza estacional">
                <b-datepicker placeholder="Fecha de vacunación Influenza estacional " icon="calendar-today"
                  :locale="locale" editable>
                </b-datepicker>
              </b-field>
            </div>
            <b-field label="EGO Y/O TIRAS URIANALISIS">
              <b-select placeholder="EGO" expanded>
                <option value="Trazas">PROT. TRAZAS</option>
                <option value="Mayor300">PROT. >300 MG/L</option>
                <option value="Leucitos">LEUCITOS >5</option>
                <option value="Nitratos">NITRATOS +</option>
                <option value="Negativo">NEGATIVO</option>
              </b-select>
            </b-field>

            <b-field label="VDRL">
              <b-select placeholder="VDRL" expanded>
                <option value="Reactivo">REACTIVO</option>
                <option value="NoReactivo">NO REACTIVO</option>
                <option value="NoRealizo">NO SE REALIZO</option>
              </b-select>
            </b-field>
            <b-field label="RESULTADO TIRAS GLUCOSA">
              <b-select placeholder="RESULTADO TIRAS GLUCOSA" expanded>
                <option value="Mayor60">>60 Y MENOR A110</option>
                <option value="Mayor110">>=110 Y MENOR O IGUAL A 126</option>
                <option value="Mayor126">>126 Y MENOR A160</option>
                <option value="NingunoGlucosa">NINGUNO</option>
              </b-select>
            </b-field>

            <b-field label="LIQUIDO AMNIOTICO">
              <b-select placeholder="LIQUIDO AMNIOTICO" expanded>
                <option value="Normal">NORMAL</option>
                <option value="Oligohidramnios">OLIGOHIDRAMNIOS</option>
                <option value="Polihidramnios">POLIHIDRAMNIOS</option>
                <option value="Anhidramnios">ANHIDRAMNIOS</option>
              </b-select>
            </b-field>
          </div>
        </div>
      </div>
    </b-step-item>

    <b-step-item step="" label="" :clickable="isStepsClickable" :type="{ 'is-success': isProfileSuccess }">
      <div>
        <div class="columns">
          <div class="column is-one-third">
            <b-field label="PLACENTA">
              <b-select placeholder="PLACENTA" expanded>
                <option value="NormalPlacen">NORMAL</option>
                <option value="Previa">PREVIA</option>
                <option value="AcretismoPlacentario">ACRETISMO PLACENTARIO</option>
                <option value="Calcificada">CALCIFICADA</option>
              </b-select>
            </b-field>


            <b-field label="¿ACUDIÓ A REFERENCIA?"></b-field>
            <b-select placeholder="REFERENCIAS" expanded>
              <option value="Hemorragia ">H.G. TLAPA</option>
              <option value="Retención de placenta">HMNIG</option>
              <option value="Placenta previa">H.C. ACATEPEC</option>
              <option value="Diabetes gestacional">H.C. ALCOZAUCA</option>
              <option value="Preeclampsia">H.C. MALINALTEPEC </option>
              <option value="Preeclampsia">H.C. OLINALÁ</option>
              <option value="Preeclampsia">H.C. TLACOAPA</option>
              <option value="Preeclampsia">H.C. XOCHIHUEHUETLÁN</option>
              <option value="Preeclampsia">H.C. ZAPOTITLÁN TABLAS</option>
              <option value="Preeclampsia">CENTRO DE SALUD</option>
              <option value="Preeclampsia">H.C. XOCHIHUEHUETLÁN</option>
              <option value="Preeclampsia">HOGAR</option>
              <option value="Preeclampsia">H.C. SAN SLUIS</option>
              <option value="Preeclampsia">H.G. AYUTLA</option>
              <option value="Preeclampsia">H.G. OMETEPEC</option>
              <option value="Preeclampsia">H.G CHUAUTLA</option>
              <option value="Preeclampsia">OTROS HOSPITALES</option>
              <option value="Preeclampsia">H.G. AYUTLA</option>
              <option value="Preeclampsia">H.G. OMETEPEC</option>
              <option value="Preeclampsia">NO ACUDIÓ</option>
            </b-select>


            <b-field label="¿SE ORIENTO EN SEÑALES DE PELIGRO DURANTE EL EMBARAZO, PARTO, PUERPERIO, RN?"></b-field>
            <b-select placeholder="¿SE ORIENTO EN SEÑALES DE PELIGRO DURANTE EL EMBARAZO?" expanded>
              <option value="Hemorragia ">SI</option>
              <option value="Retención de placenta">NO</option>
            </b-select>

            <b-field label="¿SE ORIENTO EN SEÑALES DE PELIGRO DURANTE EL EMBARAZO, PARTO, PUERPERIO, RN?"></b-field>
            <b-select placeholder="¿SE ORIENTO EN SEÑALES DE PELIGRO DURANTE EL EMBARAZO?" expanded>
              <option value="Hemorragia ">SI</option>
              <option value="Retención de placenta">NO</option>
            </b-select>
          </div>
        
        <div class="column is-one-third">
          <div class="control">
            <b-field label="Fecha probable de parto por USG">
              <b-datepicker placeholder="Fecha probable de parto por USG (1Er Trimestre)" icon="calendar-today"
                :locale="locale">
              </b-datepicker>
            </b-field>
          </div>
          <b-field label="CONTRAREFERENCIA RECIBIDA">
            <b-select placeholder="REFERENCIA" expanded>
              <option value="SiConrefe">SI</option>
              <option value="NoCorefe">NO</option>
            </b-select>
          </b-field>
          <b-field label="¿DONDE ATENDERÁN SU PARTO?">
            <b-select placeholder="¿DONDE ATENDERÁN SU PARTO?" expanded>
              <option value="SiConrefe">HOSPIAL</option>
              <option value="NoCorefe">C. SALUD</option>
              <option value="SiConrefe">HOGAR</option>
              <option value="NoCorefe">PARTICULAR
              </option>
            </b-select>
          </b-field>

          <b-field label="¿CUENTA CON TRANSPORTE AME EN CASO DE EMERGENCIA?">
            <b-select placeholder="¿CUENTA CON TRANSPORTE AME?" expanded>
              <option value="SiAME">SI</option>
              <option value="NoAME">NO</option>
            </b-select>
          </b-field>
        </div>
        <div class="column is-one-third">
          <b-field label="REFERENCIA">
            <b-select placeholder="REFERENCIA" expanded>
              <option value="SiRefe">SI</option>
              <option value="NoRefe">NO</option>
            </b-select>
          </b-field>

          <b-field label="SE LE REALIZÓ EL PLAN DE SEGURIDAD">
            <b-select placeholder="PLAN DE SEGURIDAD" expanded>
              <option value="SiRefe">SI</option>
              <option value="NoRefe">REFORZAMIENTO</option>
              <option value="NoRefe">NO</option>
            </b-select>
          </b-field>

          <b-field label="¿QUIEN ATENDERA SU PARTO?">
            <b-select placeholder="¿QUIEN ATENDERA SU PARTO?" expanded>
              <option value="SiRefe">MEDICO</option>
              <option value="NoRefe">ENFRA</option>
              <option value="NoRefe">PARTERA PROFES.</option>
              <option value="NoRefe">PARTERA TRADIC.</option>
              <option value="NoRefe">FAMILIAR</option>
              <option value="NoRefe">OTRO</option>
            </b-select>
          </b-field>
          <div class="control">
              <b-field label="Fecha de EVENTO">
                <b-datepicker placeholder="FECHA DE EVENTO" icon="calendar-today" :locale="locale"
                  editable>
                </b-datepicker>
              </b-field>
            </div>
        </div>
      </div>
      </div>
    </b-step-item>
    <b-step-item :step="showSocial ? '' : ''" label="" :clickable="isStepsClickable" disabled>
      <div class="column has-text-centered">
        <h3 class="is-size-3">Triagge</h3>
        <p>Registrar embarazada</p>

      </div>

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
    checkboxGroup: [],
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
<style>

</style>