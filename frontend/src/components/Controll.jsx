import { useEffect, useRef, useState } from "preact/hooks"

const Controll = ({ value, setVolume }) => {
    const vref = useRef()
    const [currentValue, setCurrentValue] = useState(value)
    const [startY, setStartY] = useState(null);
    const [endY, setEndY] = useState(null);
    const [percentage, setPercentage] = useState(0);

    function handleTouchStart(event) {
        setStartY(event.touches[0].clientY);
    }

    function handleTouchMove(event) {
        setEndY(event.touches[0].clientY);
        persentCounter()
    }

    function persentCounter() {
        const distance = startY - endY;
        const elementHeight = vref.current.clientHeight;
        const newPercentage = (distance / elementHeight) * 100;
        setPercentage(newPercentage);

    }
    function roundToNearest(value) {
        if (value < 50) {
          return 0;
        } else {
          return 100;
        }
      }
    const handleTouchEnd = () => {
        const val = currentValue + percentage;
        setCurrentValue(
            val>=0 && val<=100 ? val : roundToNearest(val)
        )
        setPercentage(0)
    }

    useEffect(() => {
        setVolume(currentValue);
    }, [currentValue])

    return <div className="flex-1 flex flex-col items-center justify-center">
        <div ref={vref} className="bg-slate-400 rounded-2xl flex-1 flex flex-col justify-end" onTouchStart={handleTouchStart} onTouchMove={handleTouchMove} onTouchEnd={handleTouchEnd}>
            <div
                className="w-16 flex justify-center bg-white rounded-2xl text-red-500 pt-4"
                style={{ height: (currentValue + percentage) + '%' }}
            ><div className="bg-slate-200 h-1 w-[60%] rounded-md"></div></div>
        </div>
    </div>
}

export default Controll