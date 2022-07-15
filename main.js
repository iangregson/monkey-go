
window.terminal = createTerminal();
runMonkeyRepl(window.terminal);

function createTerminal() {
  const terminalContainer = document.getElementById('terminal');

  // Clean terminal
  while (terminalContainer.children.length) {
    terminalContainer.removeChild(terminalContainer.children[0]);
  }

  const unicode11Addon = new Unicode11Addon.Unicode11Addon();
  const fitAddon = new FitAddon.FitAddon();
  const webLinksAddon = new WebLinksAddon.WebLinksAddon();

  const isWindows = ['Windows', 'Win16', 'Win32', 'WinCE'].indexOf(navigator.platform) >= 0;
  
  xterm = new Terminal({
    windowsMode: isWindows,
    cursorBlink: true,
    fontFamily: 'Fira Code, courier-new, courier, monospace'
  });

  // Load addons
  xterm.loadAddon(fitAddon);
  xterm.loadAddon(webLinksAddon);
  xterm.loadAddon(unicode11Addon);
  xterm.unicode.activeVersion = '11';

  // Load pseudoterminal
  const { master: ptyWriter, slave: ptyReader } = openpty();
  xterm.loadAddon(ptyWriter);
  
  terminalContainer.style.height = '100vh';
  xterm.open(terminalContainer);
  xterm.element.style.padding = '20px';
  
  fitAddon.fit();
  xterm.focus();
  
  ptyReader.readString = function readString() {
    const buffer = this.read();
    return utf8BytesToString(buffer)[0];
  }

  xterm.ptyReader = ptyReader;
  xterm.ptyWriter = ptyWriter;

  return xterm;
}

function runMonkeyRepl(terminal) {
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("monkey-repl.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    initMonkeyRepl(terminal, terminal.ptyReader, terminal.ptyWriter);
  });
}