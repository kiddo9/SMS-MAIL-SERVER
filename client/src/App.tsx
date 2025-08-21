import { BrowserRouter, Route, Routes } from "react-router-dom"
import Login from "./pages/auth/Login"
import NotFound from "./pages/NotFound"
import OTPVerify from "./pages/auth/OTPVerify"
import { GoogleReCaptchaProvider } from "react-google-recaptcha-v3"
import AuthContextProvider from "./contexts/AuthContext"

function App() {
  

  return (

    <BrowserRouter>
      <AuthContextProvider>
        <Routes>
          <Route path="/auth/*" element={
            <>
              <GoogleReCaptchaProvider reCaptchaKey={import.meta.env.VITE_RECAPTCHA_KEY}>
                <Routes>
                  <Route path="/login" element={<Login />} />
                  <Route path="/verify" element={<OTPVerify />} />
                  <Route path="*" element={<NotFound />} />
                </Routes>
              </GoogleReCaptchaProvider>
            </>
          } />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </AuthContextProvider>
    </BrowserRouter>
  )
}

export default App
