import { c as create_ssr_component, d as add_attribute, f as each, e as escape, v as validate_component } from "../../chunks/index.js";
function getFileByUID(uid) {
  return fetch(`/api/uid/${uid}`).then((res) => {
    return res.text();
  });
}
const FileBrowser_svelte_svelte_type_style_lang = "";
const css$2 = {
  code: "#centerer.svelte-zdnmo6.svelte-zdnmo6.svelte-zdnmo6{position:fixed;top:0;left:0;display:none;align-items:center;justify-content:center;width:100vw;height:100vh;z-index:1000;margin:auto;background-color:rgba(0, 0, 0, 0.3)	\n	}#file-search.svelte-zdnmo6.svelte-zdnmo6.svelte-zdnmo6{z-index:1000;padding:1rem;min-width:30ch;max-width:60ch;width:100%}form.svelte-zdnmo6>div.svelte-zdnmo6.svelte-zdnmo6{display:flex}.file-link-button.svelte-zdnmo6.svelte-zdnmo6.svelte-zdnmo6{border:0}.search-entry.svelte-zdnmo6.svelte-zdnmo6.svelte-zdnmo6{display:flex;text-align:center;min-height:2rem}.search-entry.svelte-zdnmo6>button.svelte-zdnmo6.svelte-zdnmo6{display:flex;text-align:center;align-items:center;width:20ch;min-width:20ch}.search-entry.svelte-zdnmo6>div.svelte-zdnmo6.svelte-zdnmo6{display:flex;font-weight:bold;flex-wrap:wrap;gap:0.25rem;margin:0.25rem}.search-entry.svelte-zdnmo6>div.svelte-zdnmo6>p.svelte-zdnmo6{display:flex;font-weight:bold}",
  map: null
};
const FileBrowser = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let nameFrag;
  let curTags;
  let centererElem;
  let initialInput;
  let results = [];
  let { selectedUID = "" } = $$props;
  let { selectedDoc = void 0 } = $$props;
  function show() {
    centererElem.style.display = "flex";
    initialInput.focus();
  }
  if ($$props.selectedUID === void 0 && $$bindings.selectedUID && selectedUID !== void 0)
    $$bindings.selectedUID(selectedUID);
  if ($$props.selectedDoc === void 0 && $$bindings.selectedDoc && selectedDoc !== void 0)
    $$bindings.selectedDoc(selectedDoc);
  if ($$props.show === void 0 && $$bindings.show && show !== void 0)
    $$bindings.show(show);
  $$result.css.add(css$2);
  return `<div id="${"centerer"}" tabindex="${"0"}" class="${"svelte-zdnmo6"}"${add_attribute("this", centererElem, 0)}><div><div id="${"file-search"}" class="${"has-background-white svelte-zdnmo6"}"><form class="${"svelte-zdnmo6"}"><label>Search by Name</label>
				<div class="${"svelte-zdnmo6"}"><input class="${"input"}"${add_attribute("this", initialInput, 0)}${add_attribute("value", nameFrag, 0)}>
					<button class="${"button has-background-primary has-text-light"}">Search</button></div></form>
			<form class="${"svelte-zdnmo6"}"><label>Search by Tags</label>
				<div class="${"svelte-zdnmo6"}"><input class="${"input"}"${add_attribute("value", curTags, 0)}>
					<button class="${"button has-background-primary has-text-light"}">Submit Tags</button></div></form>

			<ul>${each(results, (d) => {
    return `<li class="${"search-entry svelte-zdnmo6"}"><button class="${"file-link-button has-background-white svelte-zdnmo6"}">${escape(d.name)}</button>
					<div class="${"svelte-zdnmo6"}">${each(d.tags, (t) => {
      return `<p class="${"svelte-zdnmo6"}">${escape(t)}</p>`;
    })}</div>
				</li>`;
  })}</ul></div></div>
</div>`;
});
const Display_svelte_svelte_type_style_lang = "";
const css$1 = {
  code: "#display-content.svelte-h2p3mw{height:100%;width:50%;flex-grow:1;padding:0.5rem;border:solid black 1px;overflow-y:scroll\n	}",
  map: null
};
const Display = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { content = "" } = $$props;
  let contentElem;
  async function updateView() {
    console.log(content);
    contentElem.innerHTML = content;
    await MathJax.typesetPromise();
    contentElem.scrollTop = contentElem.scrollHeight;
  }
  if ($$props.content === void 0 && $$bindings.content && content !== void 0)
    $$bindings.content(content);
  $$result.css.add(css$1);
  {
    {
      if (content) {
        updateView();
      }
    }
  }
  return `<div id="${"display-content"}" class="${"content svelte-h2p3mw"}"${add_attribute("this", contentElem, 0)}></div>`;
});
const Editor_svelte_svelte_type_style_lang = "";
const css = {
  code: "#editor.svelte-18fvyn0.svelte-18fvyn0{width:100%;height:calc(80vh - 3.5rem)}#editor-content.svelte-18fvyn0.svelte-18fvyn0{display:flex;width:100%;height:80vh;flex-wrap:wrap}#editor-content.svelte-18fvyn0>.svelte-18fvyn0{min-width:30ch;width:50%;height:100%;flex-grow:1}form.svelte-18fvyn0.svelte-18fvyn0{display:flex}form.svelte-18fvyn0>div.svelte-18fvyn0{display:flex;flex-direction:column;padding:0.25rem}",
  map: null
};
const Editor = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let aceEditor;
  let { curDoc = void 0 } = $$props;
  let curName = "";
  let curTags = "";
  let displayContent;
  function setMeta() {
    if (curDoc) {
      curDoc.uid;
      curName = curDoc.name;
      curTags = curDoc.tags.join(" ");
    } else {
      curName = "";
      curTags = "";
    }
  }
  if ($$props.curDoc === void 0 && $$bindings.curDoc && curDoc !== void 0)
    $$bindings.curDoc(curDoc);
  $$result.css.add(css);
  let $$settled;
  let $$rendered;
  do {
    $$settled = true;
    {
      {
        if (curDoc) {
          getFileByUID(curDoc.uid).then((data) => {
            aceEditor.resize();
            aceEditor.setValue(data);
            aceEditor.clearSelection();
          });
        }
        setMeta();
      }
    }
    $$rendered = `<div id="${"editor-content"}" class="${"svelte-18fvyn0"}"><div style="${"border: solid black 1px;"}" class="${"svelte-18fvyn0"}"><div style="${"display: flex;"}"><form style="${"width: auto; flex-grow: 1;"}" class="${"svelte-18fvyn0"}"><div class="${"svelte-18fvyn0"}"><label>Name</label>
					<input${add_attribute("value", curName, 0)}></div>

				<div style="${"flex-grow: 1;"}" class="${"svelte-18fvyn0"}"><label>Tags</label>
					<input style="${"width: 100%;"}"${add_attribute("value", curTags, 0)}></div>


				<button style="${"display: none;"}"></button></form>

			<div style="${"display: flex; align-items: center; margin: 0em 0.25rem;"}"><button class="${"button has-background-danger has-text-light"}" style="${"font-weight: bold;"}">X</button></div></div>

		<div id="${"editor"}" class="${"svelte-18fvyn0"}"></div></div>

	${validate_component(Display, "Display").$$render(
      $$result,
      { content: displayContent },
      {
        content: ($$value) => {
          displayContent = $$value;
          $$settled = false;
        }
      },
      {}
    )}
</div>`;
  } while (!$$settled);
  return $$rendered;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let selectedDoc = void 0;
  let browserElemShow;
  let $$settled;
  let $$rendered;
  do {
    $$settled = true;
    $$rendered = `<div class="${"has-background-dark"}" style="${"display:flex; justify-content: center; align-items: center;"}"><button class="${"button has-background-primary has-text-light"}" style="${"margin-left: 0.25rem; margin-right: 0.25rem; font-weight: bold; "}">+</button>
	<button class="${"button has-background-primary has-text-light"}" style="${""}">Search</button> 
	<div style="${"flex-grow: 1;"}"></div>
	<div class="${"has-text-primary-light"}" style="${"display: flex; align-items: center; font-size: 1.5rem; font-weight: bold; text-align: center"}">SaladNotes</div>
	<div style="${"flex-grow: 1;"}"></div></div>
${validate_component(FileBrowser, "FileBrowser").$$render(
      $$result,
      { show: browserElemShow, selectedDoc },
      {
        show: ($$value) => {
          browserElemShow = $$value;
          $$settled = false;
        },
        selectedDoc: ($$value) => {
          selectedDoc = $$value;
          $$settled = false;
        }
      },
      {}
    )}
<div>${validate_component(Editor, "Editor").$$render(
      $$result,
      { curDoc: selectedDoc },
      {
        curDoc: ($$value) => {
          selectedDoc = $$value;
          $$settled = false;
        }
      },
      {}
    )}
</div>`;
  } while (!$$settled);
  return $$rendered;
});
export {
  Page as default
};
