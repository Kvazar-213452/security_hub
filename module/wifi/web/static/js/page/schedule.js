// module/wifi/web/static/js/page/schedule.js

class Chart {
    createSvgElement(tagName) {
      return document.createElementNS('http://www.w3.org/2000/svg', tagName)
    }
  
    setAttributes($svgElement, attributesObject) {
      Object.keys(attributesObject).forEach((key) => {
        $svgElement.setAttribute(key, attributesObject[key])
      })
    }
  }
  
  class LineChart extends Chart {
      horizontalPadding = 30
      legendYPadding = 30
      topYPadding = 30
      chartLineStrokeWidth = 5
      circleRadius = 0
      
      constructor(data, $container) {
          super();
          this.data = data;
          this.$container = $container;
  
          this.maxWidth = this.$container.offsetWidth;
          this.maxHeight = this.$container.offsetHeight;
  
          this.maxChartWidth = this.maxWidth - this.horizontalPadding * 3;
          this.maxChartHeight = this.maxHeight - this.legendYPadding - this.topYPadding;
  
          this.maxY = Math.max(...data.map((el) => el.y));
          this.minY = Math.min(...data.map((el) => el.y));
          this.zoom = this.maxChartHeight / (this.maxY - this.minY);
          
          if (this.zoom < 0) {
          this.zoom = 1 + this.zoom;
          }
          if (!isFinite(this.zoom)) {
          this.zoom = 1;
          }
      }
  
      createChartLine() {
          const $chartLine = this.createSvgElement('path');
          this.setAttributes($chartLine, {
          stroke: main_color,
          'stroke-width': this.chartLineStrokeWidth,
          fill: 'none',
          });
          return $chartLine;
      }
  
      createAxisXSeparator() {
          const $axisXLine = this.createSvgElement('line');
          this.setAttributes($axisXLine, {
          x1: 0,
          x2: this.maxWidth,
          y1: this.maxChartHeight + this.topYPadding + this.chartLineStrokeWidth,
          y2: this.maxChartHeight + this.topYPadding + this.chartLineStrokeWidth,
          stroke: '#12171e',
          'stroke-width': 1,
          });
          return $axisXLine;
      }
  
      createTicks() {
          const heightPerTick = 90;
          const ticksCount = this.maxChartHeight / heightPerTick;
          const tickAdd = (this.maxY - this.minY) / ticksCount;
          const $ticks = [];
          let tickValue = this.maxY;
      
          for (let i = 0; i < ticksCount; i++) {
          const currentY = heightPerTick * i + this.topYPadding - this.circleRadius;
          const $tick = this.createSvgElement('line');
          this.setAttributes($tick, {
              x1: this.horizontalPadding,
              x2: this.maxChartWidth + this.horizontalPadding,
              y1: currentY,
              y2: currentY,
              'stroke-width': 0.5,
              stroke: main_color,
          });
      
          const $text = this.createSvgElement('text');
          this.setAttributes($text, {
              x: this.maxWidth - this.horizontalPadding,
              y: currentY,
              fill: text_color,
          });
          $text.append(Math.round(tickValue));
      
          $ticks.push($tick, $text);
          tickValue -= tickAdd;
          }
          return $ticks;
      }
    
      createCircle(el, x, y) {
          const $circle = this.createSvgElement('circle');
          this.setAttributes($circle, {
              r: this.circleRadius,
              cx: x,
              cy: y,
              fill: main_color,
              stroke: main_color,
          });
          $circle.dataset.text = `${el.z}, ${el.y}`;
          $circle.classList.add('circle');
          $circle.dataset.circle = 'true';
          return $circle;
      }
  
      onCircleOver($circle) {
          const $tooltip = document.createElement('div');
          $tooltip.textContent = $circle.dataset.text;
          $tooltip.classList.add('tooltip');
          $circle.setAttribute('stroke-width', 15);
  
          Popper.createPopper($circle, $tooltip);
  
          $circle.onmouseout = () => {
              $tooltip.remove();
              $circle.setAttribute('stroke-width', 0);
              $circle.onmouseout = null;
          };
          this.$container.appendChild($tooltip);
      }
  
      create() {
          const $svg = this.createSvgElement('svg');
          this.setAttributes($svg, {
          width: '100%',
          height: '100%',
          viewBox: `0 0 ${this.maxWidth} ${this.maxHeight}`,
          });
      
          const $chartLine = this.createChartLine();
          const $ticks = this.createTicks();
          const $legendXLine = this.createAxisXSeparator();
      
          const lineLength = this.maxChartWidth / (this.data.length - 1);
          const yShift = this.minY * this.zoom;
      
          $svg.append(...$ticks, $chartLine, $legendXLine);
          let d = 'M ';
          let currentX = 0 + this.horizontalPadding;
          this.data.forEach((el, i) => {
          const x = currentX;
          const y =
              this.maxChartHeight - 
              el.y * this.zoom + 
              yShift + 
              this.topYPadding - 
              this.circleRadius;
          d += `${x} ${y} L `;
      
          const $circle = this.createCircle(el, x, y);
          const $legendXText = this.createSvgElement('text');
          this.setAttributes($legendXText, {
              x: currentX,
              y: this.maxHeight - 5,
              fill: text_color,
          });
          $legendXText.append(el.x);
      
          $svg.append($circle, $legendXText);
          currentX += lineLength;
          });
      
          d = d.slice(0, -3);
      
          $chartLine.setAttribute('d', d);
      
          this.$container.appendChild($svg);
      
          $svg.onmouseover = (e) => {
          if (e.target.dataset.circle) {
              this.onCircleOver(e.target);
          }
          };
      
          return this;
      }
  }
  
const $chartContainer = document.getElementById('chart');

let data = [];
let len = data.length;

function schedule_render(x, y) {
    y = String(y).replace('%', '');
    len += 1;

    data.push({
        x: len,
        y: y,
        z: x
    });

    
    if (data.length > 10) {
        data.shift();
    }

    if (len >= 99) {
        len = 1;
    }

    $chartContainer.innerHTML = '';
    new LineChart(data, $chartContainer).create();
}
