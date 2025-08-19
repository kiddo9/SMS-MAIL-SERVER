import { BrowserRouter, Route, Routes } from "react-router-dom"
import Login from "./pages/auth/Login"
import NotFound from "./pages/NotFound"
import OTPVerify from "./pages/auth/OTPVerify"


function App() {
  

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/auth/*" element={
          <>
            <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/verify" element={<OTPVerify />} />
              <Route path="*" element={<NotFound />} />
            </Routes>
          </>
        } />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
