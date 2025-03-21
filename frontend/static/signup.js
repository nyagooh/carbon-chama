class SignupForm extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });

    const container = document.createElement("div");
    container.classList.add("signup-container");
    const link = document.createElement("link");
    link.rel = "stylesheet";
    link.type = "text/css";
    link.href = "/static/styles.css";

    container.innerHTML = `
              <div class="left-section">
                          <h2>Capturing Moments, Creating Memories</h2>
                      </div>
                      <div class="right-section">
                          <h2>Create an account</h2>
                          <p>Already have an account? <a href="/login">Log in</a></p>
                          <form id="signup-form" action="/auth/signup" method="post">
                              <input type="text" name="username" placeholder="Username" required>
                              <input type="email" name="email" placeholder="Email" required>
                              <input type="password" name="password" placeholder="Enter your password" required>
                              <div class="terms">
                                  <input type="checkbox" id="terms" required>
                                  <label for="terms">I agree to the <a href="#">Terms & Conditions</a></label>
                              </div>
                              <button type="submit">Create Account</button>
                          </form>
                          <p>Or register with</p>
                          <button id="google-signup-btn" class="google-btn">
                          <text class = "google-text" x="10" y="30" font-family="Poppins, Arial, sans-serif"  font-weight="bold">
                                  <span class="blue">G</span>
                                  <span class="red">o</span>
                                  <span class="yellow">o</span>
                                  <span class="blue">g</span>
                                  <span class="green">l</span>
                                  <span class="red">e</span>
                          </text>
                          </button>
                      </div>
          `;

    this.shadowRoot.append(link, container);
    this.shadowRoot
      .querySelector("#google-signup-btn")
      .addEventListener("click", () => {
        window.location.href = "/auth/google";
      });
      
    // Add event listener for form submission
    this.shadowRoot.querySelector("#signup-form")
      .addEventListener("submit", (e) => {
        // Form will submit normally to the action URL
      });
  }
}

customElements.define("signup-form", SignupForm);
