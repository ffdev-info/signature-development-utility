/* variables.js */

var standardSignatureSequences = `
<div id="cloned-input-0" class="cloned-input">
    <div class="row">
      <div class="two columns">
        <label class="left-label" for="signature-input-0">Signature</label>
      </div><div class="ten columns">
        <input
          name="signature-input-0"
          class="u-full-width signature"
          type="text"
          placeholder="255044462D312E34"
          value="255044462D312E34"
          id="signature-input-0" />
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="signature-relativity-0">Anchor</label>
      </div>
      <div class="ten columns">
        <select id="signature-relativity-0" name="signature-relativity-0" class="u-full-width"
                title="Position relative to beginning or end of file">
          <option value="BOFoffset">Beginning of File</option>
          <option value="EOFoffset">End of File</option>
          <option value="VARoffset">Variable</option>
        </select>
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="offsetInput">Offset</label>
      </div><div class="ten columns">
        <input
          id="offset-0"
          name="offset-0"
          class="u-full-width"
          type="number"
          placeholder="0"
          value="0" />
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="maxOffsetInput">Max offset</label>
      </div><div class="ten columns">
        <input
          id="max-offset-0"
          name="max-offset-0"
          class="u-full-width"
          type="number"
          placeholder="0"
          value="0" />
      </div>
    </div>
</div>
`;

var newContainerFile = `
<div>
  </br>
  <div class="row">
    <div class="two columns">
      <label class="left-label" for="path-0">Path</label>
    </div><div class="ten columns">
      <input
        id="container-path"
        name="container-path",
        class="u-full-width"
        type="text"
        placeholder="path/to/file"
        value="path/to/file" />
    </div>
  </div>
</div>
`;

var newSignatureSequences = `
<div>
    <div class="row">
      <div class="two columns">
        <label class="left-label" for="signature-input">Signature</label>
      </div><div class="ten columns">
        <input
          name="PLACEHOLDER"
          class="u-full-width signature-container"
          type="text"
          placeholder="68656C6C6F20776F726C6421"
          value=""
          id="container-signature" />
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="signature-relativity">Anchor</label>
      </div><div class="ten columns">
        <select id="container-signature-relativity" name="PLACEHOLDER" class="u-full-width"
                title="Position relative to beginning or end of file">
            <option value="BOFoffset">Beginning of File</option>
            <option value="EOFoffset">End of File</option>
            <option value="VARoffset">Variable</option>
        </select>
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="offset-input">Offset</label>
      </div><div class="ten columns">
        <input
          id="container-offset"
          name="PLACEHOLDER"
          class="u-full-width"
          type="number"
          placeholder="0"
          value="0" />
      </div>
    </div>

    <div class="row">
      <div class="two columns">
        <label class="left-label" for="max-offset-input">Max offset</label>
      </div><div class="ten columns">
        <input
          id="container-max-offset"
          name="PLACEHOLDER"
          class="u-full-width"
          type="number"
          placeholder="0"
          value="0" />
      </div>
    </div>
</div>
`;
