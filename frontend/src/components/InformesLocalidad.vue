<template>
    <!-- <div class="columns"> -->
    <b-steps v-model="activeStep" :animated="isAnimated" :rounded="isRounded" :has-navigation="hasNavigation"
        :icon-prev="prevIcon" :label-position="labelPosition" :mobile-mode="mobileMode">


        <div class="column has-text-centered">
            <article class="panel is-info centered">
                <p class="panel-heading">
                    Informes
                </p>
            </article>
            <h3 class="is-size-3">Busca embarazadas por:</h3>
        </div>
        <div>
            <div class="columns">
                <div class="column">
                    <b-field label="Buscar por:">
                        <b-select placeholder="Buscar por:" expanded>
                            <option value="Localidad">Localidad</option>
                            <option value="Municipio">Municipio</option>
                            <option value="Region">Region</option>

                        </b-select>
                    </b-field>



                    <label class="radio">
                        <input type="radio" name="migrante" />
                        Migrantes
                    </label>
                    <label class="radio">
                        <input type="radio" name="mayor20" />
                        Mayor de 20 años
                        <input type="radio" name="mayor20" />
                        Menor de35 años
                    </label>


                    <b-field label="Comorbilidades">
                        <b-select placeholder="Comorbilidades" expanded>
                            <option value="Nahual">Nahual</option>
                            <option value="Mixteco">Tu´un Savi (Mixteco)</option>
                            <option value="Mephaa">Meephaa (Tlapaneco)</option>
                            <option value="Amuzgo">Amuzgo</option>
                            <option value="Ninguno">Ninguno</option>
                        </b-select>
                    </b-field>


                    <b-field label="Número de gestas">
                        <b-select placeholder="Gestas" expanded>
                            <option value="Nahual">0</option>
                            <option value="Mixteco">1-2</option>
                            <option value="Mephaa">3-5</option>
                            <option value="Amuzgo">6-8</option>
                            <option value="Ninguno">Más de 9</option>
                        </b-select>
                    </b-field>
                    <b-field label="Trimmestre">
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            1Er Trimestre
                        </label>
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            2Do Trimestre
                        </label>
                        <label class="radio">
                            <input type="radio" name="trimestre" />
                            3Er Trimestre
                        </label>
                    </b-field>
                </div>



                <div class="column">
                    <b-field label="Semana de gestión">
                        <b-select placeholder="Semana de gestión" expanded>
                            <option value="Nahual">Nahual</option>
                            <option value="Mixteco">Tu´un Savi (Mixteco)</option>
                            <option value="Mephaa">Meephaa (Tlapaneco)</option>
                            <option value="Amuzgo">Amuzgo</option>
                            <option value="Ninguno">Ninguno</option>
                        </b-select>
                    </b-field>

                    <b-field label="VIH">
                        <label class="radio">
                            <input type="radio" name="react" />
                            Reactivo
                        </label>
                        <label class="radio">
                            <input type="radio" name="react" />
                            No reactivo
                        </label>
                    </b-field>

                    <b-field label="VDRL">
                        <label class="radio">
                            <input type="radio" name="vdrl" />
                            Positivo (+)
                        </label>
                        <label class="radio">
                            <input type="radio" name="vdrl" />
                            Negativo (-)
                        </label>
                    </b-field>

                    <b-field label="Grupo RH">
                        <b-select placeholder="Grupo RH" expanded>
                            <option value="Nahual">Nahual</option>
                            <option value="Mixteco">Tu´un Savi (Mixteco)</option>
                            <option value="Mephaa">Meephaa (Tlapaneco)</option>
                            <option value="Amuzgo">Amuzgo</option>
                            <option value="Ninguno">Ninguno</option>
                        </b-select>
                    </b-field>

                    <b-field label="Referencia">
                        <label class="radio">
                            <input type="radio" name="referen" />
                            Si
                        </label>
                        <label class="radio">
                            <input type="radio" name="referen" />
                            No
                        </label>
                    </b-field>

                    <b-field label="Fecha de proximas a parir">
                        <b-datepicker placeholder="Fecha de nacimiento" icon="calendar-today" :locale="locale" editable>
                        </b-datepicker>
                    </b-field>
                    <div class="buttons">
                        <button class="button is-info">BUSCAR</button>
                    </div>
                </div>
            </div>
        </div>
    </b-steps>

</template>
<script>

</script>