import { useEffect, useState } from 'preact/hooks'
import bg from './assets/bg.jpg'
import Volume from './components/Volume'
export function App() {
  const [volume, setVolume] = useState(null)
  const [isMute, setMuteState] = useState(false);
  useEffect(()=>{
    fetch('/getVolume').then(e=>e.json()).then(e=>setVolume(e))
  },[])

  const updateVolume = (v) => {
    v = Math.round(v);
    setVolume(v)
    setMuteState(false)
    fetch('/setVolume', {method: 'POST',body: JSON.stringify({volume: v})}).then(e=>e.json).then(e=>setVolume(v))
  }
  const mute = () => fetch('/mute').then(e=>e.json).then(e=>setMuteState(true))
  const unmute = () => fetch('/unmute').then(e=>e.json).then(e=>setMuteState(false))

  return (
    <div className='min-w-screen min-h-screen bg-no-repeat bg-cover text-white' style={{backgroundImage:'url('+bg+')'}}>
     <div className="max-w-[375px] mx-auto p-[24px] flex flex-col gap-5 min-h-screen">
      <div >Your controls</div>
      <div className='text-2xl font-bold'>Change your values</div>
       {volume && <Volume currentData={volume} setVolume={updateVolume} mute={mute} unmute={unmute} isMute={isMute}/>}
     </div>
    </div>
  )
}
