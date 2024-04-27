// @ts-check

/**
 * @param {string} pin
 * @returns {Promise<boolean>}
 */
async function isPinCorrect(pin) {
  const res = await fetch("/", { method: "POST", body: pin });
  return res.status === 200;
}

/**
 * @param {boolean} isCorrect
 */
function updateResult(isCorrect) {
  const el = document.getElementById("result");
  if (!el) return;

  if (isCorrect) {
    el.textContent = "🎉 Correct! The PIN has now changed.";
    el.className = "text-success fw-semibold";
  } else {
    el.textContent = "🙅 Nope.";
    el.className = "text-danger";
  }
}

const guessForm = document.getElementById("guess-form");
if (!(guessForm instanceof HTMLFormElement))
  throw new TypeError("#guess-form was not a form");

guessForm.addEventListener("submit", async (ev) => {
  ev.preventDefault();

  const data = new FormData(guessForm);
  const guess = data.get("guess");
  if (typeof guess === "string" && guess.length === 4 && guess.match(/\d{4}/)) {
    const isCorrect = await isPinCorrect(guess);
    updateResult(isCorrect);
  }
});
