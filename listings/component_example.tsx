import [...]

interface SidebarProps {
  map: Map;
  layers: (
    | ImageLayer<ImageWMS>
    | VectorLayer<VectorSource<Feature<Geometry>>, Feature<Geometry>>
  )[];
  setLayerSidebarOpen: (isOpen: boolean) => void;
}

const SidebarLayerSelector: React.FC<SidebarProps> = ({
  map,
  layers,
  setLayerSidebarOpen,
}) => {

  [...]

  return (
    <div>
      <button
        className={`layer-sidebar-toggle-button ${
          isLayerSidebarOpen ? "open" : ""
        }`}
        onClick={toggleSidebar}>
        {"Kartvalg"}
      </button>
      <div
        className={`layer-sidebar ${isLayerSidebarOpen ? "open" : "closed"}`}>
        <h3>Kartvalg</h3>
        <ToggleLayers map={map} layers={layers} />
      </div>
    </div>
  );
};
