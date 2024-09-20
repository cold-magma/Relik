import { ThemeProvider } from '@/components/theme-provider'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from '@/components/home/home';
import Login from '@/components/login/login';
import Artifactory from '@/components/artifactory/artifactory';
import UserProfile from '@/components/user-profile/user-profile';
import Settings from '@/components/settings/settings';

function App() {

  return (
    <>
      <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
        <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/account" element={<UserProfile />} />  
            <Route path="/artifacts" element={<Artifactory />} /> 
            <Route path="/login" element={<Login />} />  
            <Route path="/settings" element={<Settings />} />  
          </Routes>
        </Router>
      </ThemeProvider>
    </>
  )
}

export default App
