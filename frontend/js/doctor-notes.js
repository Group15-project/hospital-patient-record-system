let currentPatientId = null;

document.addEventListener(
    "DOMContentLoaded",
    async () => {

        const params =
            new URLSearchParams(
                window.location.search
            );

        currentPatientId =
            params.get(
                "patientId"
            );

        if (!currentPatientId) {

            alert(
                "Patient ID missing"
            );

            return;
        }

        await loadPatient();

        await loadConsultations();

        initializeForm();
    }
);

async function loadPatient() {

    try {

        const response =
            await apiRequest(
                `/patients/${currentPatientId}`
            );

        const patient =
            response.data;

        document.getElementById(
            "patientNumber"
        ).textContent =
            patient.PatientNumber;

        document.getElementById(
            "patientName"
        ).textContent =
            `${patient.FirstName} ${patient.LastName}`;

        document.getElementById(
            "patientGender"
        ).textContent =
            patient.Gender;

        document.getElementById(
            "patientAge"
        ).textContent =
            calculateAge(
                patient.DateOfBirth
            );

    } catch (error) {

        console.error(error);
    }
}

async function loadConsultations() {

    try {

        const response =
            await apiRequest(
                `/consultations/patient/${currentPatientId}`
            );

        renderConsultations(
            response.data || []
        );

    } catch (error) {

        console.error(error);
    }
}

function renderConsultations(
    consultations
) {

    const container =
        document.getElementById(
            "notesHistory"
        );

    container.innerHTML = "";

    if (
        consultations.length === 0
    ) {

        container.innerHTML =
            `
            <div class="note-item">
                No consultation history found
            </div>
        `;

        return;
    }

    consultations.forEach(
        consultation => {

            container.innerHTML += `
                <div class="note-item">

                    <div class="note-header">

                        <h4>
                            Consultation
                        </h4>

                        <span>
                            ${new Date(
                                consultation.CreatedAt
                            ).toLocaleDateString()}
                        </span>

                    </div>

                    <p>
                        <strong>
                            Complaint:
                        </strong>
                        ${consultation.ChiefComplaint || "-"}
                    </p>

                    <p>
                        <strong>
                            Status:
                        </strong>
                        ${consultation.Status}
                    </p>

                </div>
            `;
        }
    );
}

function initializeForm() {

    const form =
        document.getElementById(
            "doctorNotesForm"
        );

    form.addEventListener(
        "submit",
        async e => {

            e.preventDefault();

            try {

                const consultation =
                    await apiRequest(
                        "/consultations",
                        "POST",
                        {
                            patient_id:
                                Number(
                                    currentPatientId
                                ),

                            chief_complaint:
                                document.getElementById(
                                    "chiefComplaint"
                                ).value
                        }
                    );

                const consultationId =
                    consultation.data.ID;

                await apiRequest(
                    "/consultations/diagnosis",
                    "POST",
                    {
                        consultation_id:
                            consultationId,

                        diagnosis:
                            document.getElementById(
                                "diagnosis"
                            ).value,

                        treatment_plan:
                            document.getElementById(
                                "treatmentPlan"
                            ).value,

                        notes:
                            `
Prescription:
${document.getElementById("prescription").value}

Follow Up:
${document.getElementById("notes").value}
`
                    }
                );

                alert(
                    "Doctor notes saved successfully"
                );

                form.reset();

                await loadConsultations();

            } catch (error) {

                console.error(error);

                alert(
                    "Failed to save notes"
                );
            }
        }
    );
}

function calculateAge(
    date
) {

    if (!date) {
        return "-";
    }

    const dob =
        new Date(date);

    const today =
        new Date();

    let age =
        today.getFullYear() -
        dob.getFullYear();

    const monthDiff =
        today.getMonth() -
        dob.getMonth();

    if (
        monthDiff < 0 ||
        (
            monthDiff === 0 &&
            today.getDate() <
            dob.getDate()
        )
    ) {
        age--;
    }

    return age;
}