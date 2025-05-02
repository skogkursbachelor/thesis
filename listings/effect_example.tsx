const [visibility, setVisibility] = useState<boolean[]>(
layers.map((layer) => layer.getVisible())
)

useEffect(() => {
layers.forEach((layer, index) => {
  if (layer.getVisible() !== visibility[index]) {
    // Use setTimeout to ensure this runs after React's render phase
    setTimeout(() => {
      layer.setVisible(visibility[index]);
    }, 0);
  }
});
}, [visibility, layers]);

